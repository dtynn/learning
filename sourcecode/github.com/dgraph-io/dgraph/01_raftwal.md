### 1. raftwal

[godoc](https://godoc.org/github.com/dgraph-io/dgraph/raftwal)



之前提到, `etcd/raft`提供了 `MemoryStorage` + `wal` 的方式 来对 raft 中的 HardState, Snapshot 和 Entry 进行持久化. `wal` 将数据直接写入文件.



而对于 dgraph 来说, 它的一个物理节点上有多个 raft group, 且 raft group 会自动新建. 此时, 所有 raft group 使用同一套底层存储会相对简单一些.

本包中, dgraph 使用 [badger](https://github.com/dgraph-io/badger) 这个同属 dgraph-io 出品的 kv 数据库来保存所有 raft group 的日志.



#### 1.1 Keys

既然不同 raft 的日志都存在同一个 kv 数据库中, 那么就需要对存储的 key 进行有效地区分.

对于一个 raft node 来说, 它通过节点 id `RaftId(uint64)` 和 组 id `gid(uint32)` 两层 来标识自己

相应地, raftwal 中的三类 key 都包含这两个 id

1. snapshotKey:

   ```go
   func (w *Wal) snapshotKey(gid uint32) []byte {
   	b := make([]byte, 14)
   	binary.BigEndian.PutUint64(b[0:8], w.id)
   	copy(b[8:10], []byte("ss"))
   	binary.BigEndian.PutUint32(b[10:14], gid)
   	return b
   }
   ```

   ​

2. hardStateKey:

   ```go
   func (w *Wal) hardStateKey(gid uint32) []byte {
   	b := make([]byte, 14)
   	binary.BigEndian.PutUint64(b[0:8], w.id)
   	copy(b[8:10], []byte("hs"))
   	binary.BigEndian.PutUint32(b[10:14], gid)
   	return b
   }
   ```

3. entryKey:

   ```go
   func (w *Wal) entryKey(gid uint32, term, idx uint64) []byte {
   	b := make([]byte, 28)
   	binary.BigEndian.PutUint64(b[0:8], w.id)
   	binary.BigEndian.PutUint32(b[8:12], gid)
   	binary.BigEndian.PutUint64(b[12:20], term)
   	binary.BigEndian.PutUint64(b[20:28], idx)
   	return b
   }
   ```



#### 1.2 Wal

Wal 提供 raft 数据的读写.

对于 raft 数据的持久化, 最重要的是保证数据的一致性.



##### StoreSnapshot

```go
func (w *Wal) StoreSnapshot(gid uint32, s raftpb.Snapshot) error {
	txn := w.wals.NewTransactionAt(1, true)
	defer txn.Discard()
	
	// ...
	
	if err := txn.Set(w.snapshotKey(gid), data); err != nil {
		return err
	}
	
	// ...
	
	// 清除 snapshot 数据之前的所有 entry
	// Delete all entries before this snapshot to save disk space.
	start := w.entryKey(gid, 0, 0)
	last := w.entryKey(gid, s.Metadata.Term, s.Metadata.Index)
	
	// 这里利用了 badger 的特性, 在遍历的时候仅读取 key 数据, 减少了读取 value 带来的开销
	opt := badger.DefaultIteratorOptions
	opt.PrefetchValues = false
	itr := txn.NewIterator(opt)
	defer itr.Close()

	// 逐一删除不再需要的 entry
	for itr.Seek(start); itr.Valid(); itr.Next() {
		// ...
	}

	// Failure to delete entries is not a fatal error, so should be
	// ok to ignore
	if err := txn.CommitAt(1, nil); err != nil {
		x.Printf("Error while storing snapshot %v\n", err)
		return err
	}
	return nil
}
```





##### Store

```go
// Store stores the hardstate and entries for a given RAFT group.
func (w *Wal) Store(gid uint32, h raftpb.HardState, es []raftpb.Entry) error {
	txn := w.wals.NewTransactionAt(1, true)

	var t, i uint64
	// 逐一保存 entry
	for _, e := range es {
		t, i = e.Term, e.Index
		
		// ...
	}

	// 如果有必要, 保存 HardState
	if !raft.IsEmptyHardState(h) {
		// ...
	}

	// If we get no entries, then the default value of t and i would be zero. That would
	// end up deleting all the previous valid raft entry logs. This check avoids that.
	if t > 0 || i > 0 {
		// When writing an Entry with Index i, any previously-persisted entries
		// with Index >= i must be discarded.
		// Ideally we should be deleting entries from previous term with index >= i,
		// but to avoid complexity we remove them during reading from wal.
		// 有可能出现某个时间点之后, 由于网络原因, 数据分叉的情形.
		// 为了在网络恢复之后保证数据一致性, 对于每一批 entry, 需要清除逻辑上排在这批数据之后的 entry.
		start := w.entryKey(gid, t, i+1)
		prefix := w.prefix(gid)
		// ...
		
		// 逐一清除
		for itr.Seek(start); itr.ValidForPrefix(prefix); itr.Next() {
			// ...
		}
	}
	if err := txn.CommitAt(1, nil); err != nil {
		return err
	}
	return nil
}
```



##### 读取

```go
func (w *Wal) Snapshot(gid uint32) (snap raftpb.Snapshot, rerr error) {
	// ...
}
```



```go
func (w *Wal) HardState(gid uint32) (hd raftpb.HardState, rerr error) {
	// ...
}
```



```go
func (w *Wal) Entries(gid uint32, fromTerm, fromIndex uint64) (es []raftpb.Entry, rerr error) {
	// ...
}
```



#### 1.3 关于badger

badger 来源于这篇论文 [WiscKey: Separating Keys from Values in SSD-conscious Storage.](https://www.usenix.org/system/files/conference/fast16/fast16-papers-lu.pdf) .

知乎上仅有的评论里, 对它的评价不甚高 [如何评价 Badger (fast key-value storage)](https://www.zhihu.com/question/59895275/answer/170359113) 😂.

但不论怎样, 它在一些情况下确实比较 *快*, 也可能非常适合 dgraph 的使用场景.

