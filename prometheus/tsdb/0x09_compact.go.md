#### Compact.go

对底层存储的压缩相关的实现

```go
// Compactor provides compaction against an underlying storage
// of time series data.
type Compactor interface {
	// Plan returns a set of non-overlapping directories that can
	// be compacted concurrently.
	// Results returned when compactions are in progress are undefined.
	Plan(dir string) ([]string, error)

	// Write persists a Block into a directory.
	Write(dest string, b BlockReader, mint, maxt int64) (ulid.ULID, error)

	// Compact runs compaction against the provided directories. Must
	// only be called concurrently with results of Plan().
	Compact(dest string, dirs ...string) (ulid.ULID, error)
}
```



##### LeveledCompactor

是 Compactor 的实现

###### Plan

```go
// Plan returns a list of compactable blocks in the provided directory.
func (c *LeveledCompactor) Plan(dir string) ([]string, error) {
	dirs, err := blockDirs(dir)
	
	// ...
  
	var dms []dirMeta

	for _, dir := range dirs {
		// 读取 BlockMeta 作为判断是否可以 compact 的依据
		meta, err := readMetaFile(dir)
		
		// ...
	}
	return c.plan(dms)
}
```



###### populateBlock

`LeveledCompactor.Write` 和 `LeveledCompactor.Compact` 两个方法中都用到 `LeveledCompactor.write`, 而 `LeveledCompactor.populateBlock` 是 write 方法的重要逻辑.

其作用是将一组 Block 的数据合并, 再写入 IndexWriter, ChunkWriter.

```go
// populateBlock fills the index and chunk writers with new data gathered as the union
// of the provided blocks. It returns meta information for the new block.
func (c *LeveledCompactor) populateBlock(blocks []BlockReader, meta *BlockMeta, indexw IndexWriter, chunkw ChunkWriter) error {
	var (
		set        ChunkSeriesSet
		allSymbols = make(map[string]struct{}, 1<<16)
		closers    = []io.Closer{}
	)
	defer func() { closeAll(closers...) }()

  	// 遍历旧 block 数据
	for i, b := range blocks {
		indexr, err := b.Index()
		// ...

		chunkr, err := b.Chunks()
		// ...

		tombsr, err := b.Tombstones()
		// ...

		symbols, err := indexr.Symbols()
		// ...

		all, err := indexr.Postings(index.AllPostingsKey())
		if err != nil {
			return err
		}
		all = indexr.SortedPostings(all)

		s := newCompactionSeriesSet(indexr, chunkr, tombsr, all)

		// ...
      
		// 与上一层并形成一个新的 merger
		set, err = newCompactionMerger(set, s)
		if err != nil {
			return err
		}
	}

	// We fully rebuild the postings list index from merged series.
	// ...

	// 遍历 merger
	for set.Next() {
		lset, chks, dranges := set.At() // The chunks here are not fully deleted.

		// Skip the series with all deleted chunks.
		// ...

		if err := chunkw.WriteChunks(chks...); err != nil {
			return errors.Wrap(err, "write chunks")
		}

		if err := indexw.AddSeries(i, lset, chks...); err != nil {
			return errors.Wrap(err, "add series")
		}

		// ...
	}
	
	// ...

	s := make([]string, 0, 256)
	for n, v := range values {
		// ...

		if err := indexw.WriteLabelIndex([]string{n}, s); err != nil {
			return errors.Wrap(err, "write label index")
		}
	}

	for _, l := range postings.SortedKeys() {
		if err := indexw.WritePostings(l.Name, l.Value, postings.Get(l.Name, l.Value)); err != nil {
			return errors.Wrap(err, "write postings")
		}
	}
	return nil
}
```



