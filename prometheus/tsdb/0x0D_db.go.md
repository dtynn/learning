#### db.go



##### Appender

Appender 是写入接口, *Head 就实现了 Appender

```
// Appender allows appending a batch of data. It must be completed with a
// call to Commit or Rollback and must not be reused afterwards.
//
// Operations on the Appender interface are not goroutine-safe.
type Appender interface {
	// Add adds a sample pair for the given series. A reference number is
	// returned which can be used to add further samples in the same or later
	// transactions.
	// Returned reference numbers are ephemeral and may be rejected in calls
	// to AddFast() at any point. Adding the sample via Add() returns a new
	// reference number.
	// If the reference is the empty string it must not be used for caching.
	Add(l labels.Labels, t int64, v float64) (uint64, error)

	// Add adds a sample pair for the referenced series. It is generally faster
	// than adding a sample by providing its full label set.
	AddFast(ref uint64, t int64, v float64) error

	// Commit submits the collected samples and purges the batch.
	Commit() error

	// Rollback rolls back all modifications made in the appender so far.
	Rollback() error
}
```





##### DB

DB 是向调用者提供的最主要的结构体.

```go
// DB handles reads and writes of time series falling into
// a hashed partition of a seriedb.
type DB struct {
	dir   string
	lockf *lockfile.Lockfile

	logger    log.Logger
	metrics   *dbMetrics
	opts      *Options
	chunkPool chunkenc.Pool
	compactor Compactor

	// Mutex for that must be held when modifying the general block layout.
	mtx    sync.RWMutex
	blocks []*Block

	head *Head

	compactc chan struct{}
	donec    chan struct{}
	stopc    chan struct{}

	// cmtx is used to control compactions and deletions.
	cmtx               sync.Mutex
	compactionsEnabled bool
}
```



###### reload

```go
// reload on-disk blocks and trigger head truncation if new blocks appeared. It takes
// a list of block directories which should be deleted during reload.
func (db *DB) reload(deleteable ...string) (err error) {
	// ...
	
	// 读取当前所有的 block 目录
	dirs, err := blockDirs(db.dir)
	
	// ...
	
	var (
		blocks []*Block
		exist  = map[ulid.ULID]struct{}{}
	)

	for _, dir := range dirs {
		meta, err := readMetaFile(dir)
		
		// ...

		// 尝试获取目录对应的 Block, 先从内存, 再从硬盘
		b, ok := db.getBlock(meta.ULID)
		if !ok {
			b, err = OpenBlock(dir, db.chunkPool)
			
			// ...
		}

		blocks = append(blocks, b)
		exist[meta.ULID] = struct{}{}
	}

	// 按照 Block 覆盖的时间重新排序
	if err := validateBlockSequence(blocks); err != nil {
		return errors.Wrap(err, "invalid block sequence")
	}

	// ...
	
	// 清除不必要的 Block 文件
	for _, b := range oldBlocks {
		if _, ok := exist[b.Meta().ULID]; ok {
			continue
		}
		if err := b.Close(); err != nil {
			level.Warn(db.logger).Log("msg", "closing block failed", "err", err)
		}
		if err := os.RemoveAll(b.Dir()); err != nil {
			level.Warn(db.logger).Log("msg", "deleting block failed", "err", err)
		}
	}

	// Garbage collect data in the head if the most recent persisted block
	// covers data of its current time range.
	if len(blocks) == 0 {
		return nil
	}
	maxt := blocks[len(blocks)-1].Meta().MaxTime

	return errors.Wrap(db.head.Truncate(maxt), "head truncate failed")
}
```



###### run

run 方法在 Open 时被调用, 在一个单独的 goroutine 中执行, 主要是定期对数据进行压缩以节省空间

```go
func (db *DB) run() {
	defer close(db.donec)

	backoff := time.Duration(0)

	for {
		select {
		case <-db.stopc:
			return
		case <-time.After(backoff):
		}

		select {
		case <-time.After(1 * time.Minute):
			select {
			case db.compactc <- struct{}{}:
			default:
			}
		case <-db.compactc:
			// 执行压缩相关代码

		case <-db.stopc:
			return
		}
	}
}
```



###### Appender

返回的是封装的结果 dbAppender, 后面专门再分析



###### Qurier

返回的是所有指定时间范围内的 Block 聚合

```
// Querier returns a new querier over the data partition for the given time range.
// A goroutine must not handle more than one open Querier.
func (db *DB) Querier(mint, maxt int64) (Querier, error) {
	var blocks []BlockReader

	db.mtx.RLock()
	defer db.mtx.RUnlock()

	for _, b := range db.blocks {
		m := b.Meta()
		
		// 找出符合时间段的 block
		if intervalOverlap(mint, maxt, m.MinTime, m.MaxTime) {
			blocks = append(blocks, b)
		}
	}
	
	// 前面提到, Head 可以视作当前 Block
	if maxt >= db.head.MinTime() {
		blocks = append(blocks, db.head)
	}

	// Block 的聚合
	sq := &querier{
		blocks: make([]Querier, 0, len(blocks)),
	}
	for _, b := range blocks {
		q, err := NewBlockQuerier(b, mint, maxt)
		if err == nil {
			sq.blocks = append(sq.blocks, q)
			continue
		}
		// If we fail, all previously opened queriers must be closed.
		for _, q := range sq.blocks {
			q.Close()
		}
		return nil, errors.Wrapf(err, "open querier for block %s", b)
	}
	return sq, nil
}
```



###### Delete

这边实际会将 Delete 操作分给各个受影响的 Block



###### CleanTombstone

前面提到, 各个 Block Delete 内的逻辑实际是写 WAL 以及 Tombstone 文件

这里会对当前所有 Block 真正进行清理, 然后调用 `reload` 方法.



##### dbAppender

是对 *headAppender 的封装, 在 Commit 的时候触发 compact

```
// Appender opens a new appender against the database.
func (db *DB) Appender() Appender {
	return dbAppender{db: db, Appender: db.head.Appender()}
}

// dbAppender wraps the DB's head appender and triggers compactions on commit
// if necessary.
type dbAppender struct {
	Appender
	db *DB
}

func (a dbAppender) Commit() error {
	err := a.Appender.Commit()

	// We could just run this check every few minutes practically. But for benchmarks
	// and high frequency use cases this is the safer way.
	if a.db.head.MaxTime()-a.db.head.MinTime() > a.db.head.chunkRange/2*3 {
		select {
		case a.db.compactc <- struct{}{}:
		default:
		}
	}
	return err
}
```

