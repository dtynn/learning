#### wal.go

prometheus/tsdb 会将几类数据先写入 wal (write ahead log) 文件

```go
// WALEntryType indicates what data a WAL entry contains.
type WALEntryType uint8

// Entry types in a segment file.
const (
	WALEntrySymbols WALEntryType = 1
	WALEntrySeries  WALEntryType = 2
	WALEntrySamples WALEntryType = 3
	WALEntryDeletes WALEntryType = 4
)
```



```go
// WAL is a write ahead log that can log new series labels and samples.
// It must be completely read before new entries are logged.
type WAL interface {
	Reader() WALReader
	LogSeries([]RefSeries) error
	LogSamples([]RefSample) error
	LogDeletes([]Stone) error
	Truncate(mint int64, keep func(uint64) bool) error
	Close() error
}

// WALReader reads entries from a WAL.
type WALReader interface {
	Read(
		seriesf func([]RefSeries),
		samplesf func([]RefSample),
		deletesf func([]Stone),
	) error
}
```



与之相关的数据结构定义如下

```
// RefSeries is the series labels with the series ID.
type RefSeries struct {
	Ref    uint64
	Labels labels.Labels
}

// RefSample is a timestamp/value pair associated with a reference to a series.
type RefSample struct {
	Ref uint64
	T   int64
	V   float64

	// 基于内存的 series 数据, 在后续的阅读中再仔细分析
	series *memSeries
}

```



##### SegmentWAL

这是 WAL 的一个实现, 会将数据切成 256MB 一片进行存储, 切片的组织方式与 chunks 类似.

相应的, 操作文件的相关实现代码也很相似.

```
// segmentFile wraps a file object of a segment and tracks the highest timestamp
// it contains. During WAL truncating, all segments with no higher timestamp than
// the truncation threshold can be compacted.
type segmentFile struct {
	*os.File
	maxTime   int64  // highest tombstone or sample timpstamp in segment
	minSeries uint64 // lowerst series ID in segment
}

// SegmentWAL is a write ahead log for series data.
type SegmentWAL struct {
	mtx     sync.Mutex
	metrics *walMetrics

	dirFile *os.File
	files   []*segmentFile

	logger        log.Logger
	flushInterval time.Duration
	segmentSize   int64

	crc32 hash.Hash32
	cur   *bufio.Writer
	curN  int64

	// 信号
	stopc   chan struct{}
	donec   chan struct{}
	
	// 后台执行的操作
	actorc  chan func() error // sequentialized background operations
	
	buffers sync.Pool
}
```



###### LogXXXX

LogSeries, LogSamples, LogDeletes 对各自的操作数据分别编码写入 WAL.



###### Truncate

```go
// Truncate deletes the values prior to mint and the series which the keep function
// does not indiciate to preserve.
// 用于清除不再需要的数据
func (w *SegmentWAL) Truncate(mint int64, keep func(uint64) bool) error {
	// ...

	return nil
}
```



###### run

通过 `OpenSegmentWAL` 打开一个 SegmentWAL 的时候, 会在一个独立的 goroutine 中运行 run 函数, 用来处理 `actorc` 传递的后台操作.

目前 `actorc` 传递的操作仅有文件的分片

```go
// cut finishes the currently active segments and opens the next one.
// The encoder is reset to point to the new segment.
func (w *SegmentWAL) cut() error {
	// Sync current head to disk and close.
	if hf := w.head(); hf != nil {
		if err := w.flush(); err != nil {
			return err
		}
		
		// Finish last segment asynchronously to not block the WAL moving along
		// in the new segment.
		// 结束当前的切片文件
		go func() {
			w.actorc <- func() error {
				off, err := hf.Seek(0, os.SEEK_CUR)
				if err != nil {
					return errors.Wrapf(err, "finish old segment %s", hf.Name())
				}
				if err := hf.Truncate(off); err != nil {
					return errors.Wrapf(err, "finish old segment %s", hf.Name())
				}
				if err := hf.Sync(); err != nil {
					return errors.Wrapf(err, "finish old segment %s", hf.Name())
				}
				if err := hf.Close(); err != nil {
					return errors.Wrapf(err, "finish old segment %s", hf.Name())
				}
				return nil
			}
		}()
	}

	// 初始化新的切片文件供写入
	// ...
	
	return nil
}
```



