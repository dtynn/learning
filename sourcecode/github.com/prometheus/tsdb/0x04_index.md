#### index

```
index
├── encoding_helpers.go
├── index.go
├── index_test.go
├── postings.go
└── postings_test.go
```

[index](https://godoc.org/github.com/prometheus/tsdb/index) 实现针对 labels 的索引.

在 prometheus/tsdb 中, 认为 labels + timestamp + value 是一个完整的数据点

chunks 相关的代码用于存储 timestamp + value, 而 index 则是对于 labels 的处理.



##### encbuf, decbuf

作为 buffer 在 index 数据编码/解码时进行 复用

实际上这里也定义了一些数据格式如何进行存储

```go
// putVarintStr writes a string to the buffer prefixed by its varint length (in bytes!).
// 对于字符串, 分别写入长度及字符串本身
func (e *encbuf) putUvarintStr(s string) {
	b := *(*[]byte)(unsafe.Pointer(&s))
	e.putUvarint(len(b))
	e.putString(s)
}

// 相应地, 在解码时也会先确定 str 长度, 再从整个 []byte 中取出必要的部分
func (d *decbuf) uvarintStr() string {
	l := d.uvarint64()
	if d.e != nil {
		return ""
	}
	if len(d.b) < int(l) {
		d.e = errInvalidSize
		return ""
	}
	s := string(d.b[:l])
	d.b = d.b[l:]
	return s
}
```



##### indexWriterSeries

```go
type indexWriterSeries struct {
	// labels 的实际内容, 即 kv 对
	labels labels.Labels
	
	// 这里重要的实际是 Meta.Ref, 即每个 chunk 对应的文件/起点
	chunks []chunks.Meta // series file offset of chunks
	
	// 这里是 labels 数据在文件中的 offset
	offset uint32        // index file offset of series reference
}
```



##### indexTOC

index table of contents, 记录 index 不同类型数据的位置

```go
type indexTOC struct {
	symbols           uint64
	series            uint64
	labelIndices      uint64
	labelIndicesTable uint64
	postings          uint64
	postingsTable     uint64
}
```



##### Writer

实现 [IndexWriter](https://github.com/prometheus/tsdb/blob/d45b595a1daefad23c09a2d994bf956f8b5f15a9/block.go#L32-L58) , 基于文件的 index 存储

index 的文件格式要比 chunk 复杂的多, 可以参考 [Documentation/format/index.md](https://github.com/prometheus/tsdb/blob/d45b595a1daefad23c09a2d994bf956f8b5f15a9/Documentation/format/index.md)



每个 index 文件的写入分为 5 个阶段, 顺序执行.

```go
type indexWriterStage uint8

const (
	idxStageNone indexWriterStage = iota
	idxStageSymbols
	idxStageSeries
	idxStageLabelIndex
	idxStagePostings
	idxStageDone
)

// ensureStage handles transitions between write stages and ensures that IndexWriter
// methods are called in an order valid for the implementation.
func (w *Writer) ensureStage(s indexWriterStage) error {
	if w.stage == s {
		return nil
	}
	
	// 排在当前阶段之前的, 不可再执行
	if w.stage > s {
		return errors.Errorf("invalid stage %q, currently at %q", s, w.stage)
	}

	// Mark start of sections in table of contents.
	switch s {
	
	// ...

	// 执行到完成阶段时, 自动写入必要的辅助信息
	case idxStageDone:
		w.toc.labelIndicesTable = w.pos
		if err := w.writeOffsetTable(w.labelIndexes); err != nil {
			return err
		}
		w.toc.postingsTable = w.pos
		if err := w.writeOffsetTable(w.postings); err != nil {
			return err
		}
		if err := w.writeTOC(); err != nil {
			return err
		}
	}

	w.stage = s
	return nil
}
```



###### AddSymbols

```go
func (w *Writer) AddSymbols(sym map[string]struct{}) error {
	if err := w.ensureStage(idxStageSymbols); err != nil {
		return err
	}
	
	// ...
	
	return errors.Wrap(err, "write symbols")
}
```



label 中的每一个键或值都是一个 symbol.

通过 "使用对 symbol 的引用" 的方式, 来缩减后续索引文件中的空间占用.



###### AddSeries

```go
func (w *Writer) AddSeries(ref uint64, lset labels.Labels, chunks ...chunks.Meta) error {
	if err := w.ensureStage(idxStageSeries); err != nil {
		return err
	}
	if labels.Compare(lset, w.lastSeries) <= 0 {
		return errors.Errorf("out-of-order series added with label set %q", lset)
	}

	// 记录每个时间序列的位置
	if _, ok := w.seriesOffsets[ref]; ok {
		return errors.Errorf("series with reference %d already added", ref)
	}
	w.seriesOffsets[ref] = w.pos

	w.buf2.reset()
	w.buf2.putUvarint(len(lset))

	// 对于每个 label, 分别记录 它的 name 和 value 在索引文件中的位置
	for _, l := range lset {
		offset, ok := w.symbols[l.Name]
		// ...

		offset, ok = w.symbols[l.Value]
		// ...
	}

	w.buf2.putUvarint(len(chunks))

	// 对于 chunk 数据, 记录它覆盖的时间范围, 以及存储地址
	// 除第一个 chunk 外, 其他记录的都是变化量
	if len(chunks) > 0 {
		c := chunks[0]
		w.buf2.putVarint64(c.MinTime)
		w.buf2.putUvarint64(uint64(c.MaxTime - c.MinTime))
		w.buf2.putUvarint64(c.Ref)
		t0 := c.MaxTime
		ref0 := int64(c.Ref)

		for _, c := range chunks[1:] {
			w.buf2.putUvarint64(uint64(c.MinTime - t0))
			w.buf2.putUvarint64(uint64(c.MaxTime - c.MinTime))
			t0 = c.MaxTime

			w.buf2.putVarint64(int64(c.Ref) - ref0)
			ref0 = int64(c.Ref)
		}
	}

	// ...

	return nil
}
```



###### WriteLabelIndex

```go
// 这里传入的参数可以认为是下述结构
// 其中每一组 value 都是 names 的一组取值组合
// type Label struct {
// 	names []string
// 	valus [][]string
// }
func (w *Writer) WriteLabelIndex(names []string, values []string) error {
	if len(values)%len(names) != 0 {
		return errors.Errorf("invalid value list length %d for %d names", len(values), len(names))
	}
	if err := w.ensureStage(idxStageLabelIndex); err != nil {
		return errors.Wrap(err, "ensure stage")
	}

	// ...

	// 所有 hash entry 会统一在后续阶段写入
	w.labelIndexes = append(w.labelIndexes, hashEntry{
		keys:   names,
		offset: w.pos,
	})

	// ...

	// 对于每个 value, 都只写入引用值
	for _, v := range valt.s {
		offset, ok := w.symbols[v]
		// ...
	}

	// ...

	err = w.write(w.buf1.get(), w.buf2.get())
	return errors.Wrap(err, "write label index")
}
```



###### WritePostings

```go
// Postings 用来记录每一个 label (一对 name, value) 对应了哪些数据块, 用于检索
func (w *Writer) WritePostings(name, value string, it Postings) error {
	// ...
	
	// 每一对 name-value 对应的数据位置
	w.postings = append(w.postings, hashEntry{
		keys:   []string{name, value},
		offset: w.pos,
	})

	// Order of the references in the postings list does not imply order
	// of the series references within the persisted block they are mapped to.
	// We have to sort the new references again.
	refs := w.uint32s[:0]

	for it.Next() {
		offset, ok := w.seriesOffsets[it.At()]
		
		// ...
		
		refs = append(refs, uint32(offset))
	}
	if err := it.Err(); err != nil {
		return err
	}
	sort.Sort(uint32slice(refs))

	// ...

	err := w.write(w.buf1.get(), w.buf2.get())
	return errors.Wrap(err, "write postings")
}
```



###### Close

```go
func (w *Writer) Close() error {
	// 这里会自动执行 labelIndexes, postings, toc 的写入
	if err := w.ensureStage(idxStageDone); err != nil {
		return err
	}
	
	// 文件落盘
	if err := w.fbuf.Flush(); err != nil {
		return err
	}
	if err := fileutil.Fsync(w.f); err != nil {
		return err
	}
	return w.f.Close()
}
```



```go
// writeOffsetTable writes a sequence of readable hash entries.
func (w *Writer) writeOffsetTable(entries []hashEntry) error {
	w.buf2.reset()
	w.buf2.putBE32int(len(entries))

	for _, e := range entries {
		w.buf2.putUvarint(len(e.keys))
		for _, k := range e.keys {
			w.buf2.putUvarintStr(k)
		}
		w.buf2.putUvarint64(e.offset)
	}

	w.buf1.reset()
	w.buf1.putBE32int(w.buf2.len())
	w.buf2.putHash(w.crc32)

	return w.write(w.buf1.get(), w.buf2.get())
}
```



```go
func (w *Writer) writeTOC() error {
	w.buf1.reset()

	w.buf1.putBE64(w.toc.symbols)
	w.buf1.putBE64(w.toc.series)
	w.buf1.putBE64(w.toc.labelIndices)
	w.buf1.putBE64(w.toc.labelIndicesTable)
	w.buf1.putBE64(w.toc.postings)
	w.buf1.putBE64(w.toc.postingsTable)

	w.buf1.putHash(w.crc32)

	return w.write(w.buf1.get())
}
```



##### Reader

实现了 [IndexReader](https://github.com/prometheus/tsdb/blob/d45b595a1daefad23c09a2d994bf956f8b5f15a9/block.go#L61-L89)



```go
func newReader(b ByteSlice, c io.Closer) (*Reader, error) {
	r := &Reader{
		// ...
	}
	// Verify magic number.
	
	// ...

  	// toc 在文件尾部, 且长度固定, 因此可以直接读出
	if err := r.readTOC(); err != nil {
		return nil, errors.Wrap(err, "read TOC")
	}
	if err := r.readSymbols(int(r.toc.symbols)); err != nil {
		return nil, errors.Wrap(err, "read symbols")
	}
	var err error

	err = r.readOffsetTable(r.toc.labelIndicesTable, func(key []string, off uint32) error {
		// 不知道这里为什么会强制长度为 1?
		// 根据 Writer.WriteLabelIndex 的定义, 明显是支持多 names 的
		// 实际验证, 多 names 写入没有问题, 但在读取的时候会在这里报错
		// 等待后续看相关代码来理解吧.
		if len(key) != 1 {
			return errors.Errorf("unexpected key length %d", len(key))
		}
		r.labels[key[0]] = off
		return nil
	})
	if err != nil {
		return nil, errors.Wrap(err, "read label index table")
	}
	err = r.readOffsetTable(r.toc.postingsTable, func(key []string, off uint32) error {
		// ...
		return nil
	})
	if err != nil {
		return nil, errors.Wrap(err, "read postings table")
	}

	r.dec = &DecoderV1{symbols: r.symbols}

	return r, nil
}
```



##### Postings

Posting 及其实现的具体作用, 待阅读剩余部分的代码后再回过头来确认.

这是一个 Iterator.

```go
// Postings provides iterative access over a postings list.
type Postings interface {
	// Next advances the iterator and returns true if another value was found.
	Next() bool

	// Seek advances the iterator to value v or greater and returns
	// true if a value was found.
	Seek(v uint64) bool

	// At returns the value at the current iterator position.
	At() uint64

	// Err returns the last error of the iterator.
	Err() error
}
```



- 给出了 Posting 的交集, 并集, 以及差集实现

  ```go
  // Intersect returns a new postings list over the intersection of the
  // input postings.
  func Intersect(its ...Postings) Postings {
  	if len(its) == 0 {
  		return emptyPostings
  	}
  	if len(its) == 1 {
  		return its[0]
  	}
  	l := len(its) / 2
  	return newIntersectPostings(Intersect(its[:l]...), Intersect(its[l:]...))
  }

  type intersectPostings struct {
  	a, b     Postings
  	aok, bok bool
  	cur      uint64
  }


  ```

  ```go
  // Merge returns a new iterator over the union of the input iterators.
  func Merge(its ...Postings) Postings {
  	if len(its) == 0 {
  		return EmptyPostings()
  	}
  	if len(its) == 1 {
  		return its[0]
  	}
  	l := len(its) / 2
  	return newMergedPostings(Merge(its[:l]...), Merge(its[l:]...))
  }

  type mergedPostings struct {
  	a, b        Postings
  	initialized bool
  	aok, bok    bool
  	cur         uint64
  }


  ```

  ```go
  // Without returns a new postings list that contains all elements from the full list that
  // are not in the drop list
  func Without(full, drop Postings) Postings {
  	return newRemovedPostings(full, drop)
  }

  type removedPostings struct {
  	full, remove Postings

  	cur uint64

  	initialized bool
  	fok, rok    bool
  }
  ```

  ​


- 给出了几种特定类型的 Postings

  ```go
  // EmptyPostings returns a postings list that's always empty.
  func EmptyPostings() Postings {
  	return emptyPostings
  }

  ```

  ```go
  // ErrPostings returns new postings that immediately error.
  func ErrPostings(err error) Postings {
  	return errPostings{err}
  }


  ```

  ```go
  // listPostings implements the Postings interface over a plain list.
  type listPostings struct {
  	list []uint64
  	cur  uint64
  }


  ```

  ```go
  // bigEndianPostings implements the Postings interface over a byte stream of
  // big endian numbers.
  type bigEndianPostings struct {
  	list []byte
  	cur  uint32
  }
  ```

  ​

##### MemPostings

label - posting idx 的映射记录器

```go
// MemPostings holds postings list for series ID per label pair. They may be written
// to out of order.
// ensureOrder() must be called once before any reads are done. This allows for quick
// unordered batch fills on startup.
type MemPostings struct {
	mtx     sync.RWMutex
  
	// label 和 posting id 的关联
	m       map[labels.Label][]uint64
  
	// 成功执行 EnsureOrder 之后置为 true
	ordered bool
}
```

