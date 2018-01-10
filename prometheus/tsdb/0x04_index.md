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

```
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

```
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

index table of contents, 记录 index 各个 stage 的位置

```
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

 