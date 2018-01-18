### 0. Intro

#### 0.1 raft

[raft](https://raft.github.io/) 是一种分布式一致性算法. 

简单来说, raft 的使用场景是 log replication.

关于分布式一致性算法, paxos, raft 等概念, 网上有大量的文章, 这里不再多做说明.



##### raft in go

比较成熟的 raft golang 实现有 [hashicorp/raft](https://github.com/hashicorp/raft) 和 [etcd/raft](https://github.com/coreos/etcd/tree/master/raft) 两个版本.

二者分别被用在 [hashicorp/consul](https://github.com/hashicorp/consul) 和 [coreos/etcd](https://github.com/coreos/etcd/tree/master/raft) 中, 均有大量生产环境的使用案例.

有很多优秀的开源项目使用二者之一, 比如

- hashicorp/raft:
  - [ipfs/ipfs-cluster](https://github.com/ipfs/ipfs-cluster)
  - [rqlite/rqlite](https://github.com/rqlite/rqlite)
- etcd/raft:
  - [dgraph-io/dgraph](https://github.com/dgraph-io/dgraph)
  - [cockroachdb/cockroach](https://github.com/cockroachdb/cockroach/)



#### 0.2 etcd/raft

相比而言, `hashicorp/raft` 比较容易上手; 而 `etcd/raft` 则基于简洁的抽象, 提供了更多可能性. 

很多 `etcd/raft` 的使用者选择自行实现高度定制的网络传输层和持久化层等组件.

etcd/raft README 的 Usage 段落用了很大篇幅来描述`user has a few responsibilities`, 这些定制化的组件也就是在这里做文章

> 可以参考 cockroach 的开发博客 [Scaling Raft](https://www.cockroachlabs.com/blog/scaling-raft/) . (文中 Multi-Raft 的链接已经失效, 这是因为 cockroach 的开发者发现这套实现很难从使用应用中解耦出来 [etcd/issues/4932](https://github.com/coreos/etcd/issues/4932) )
>
> TiDB 开发过程中遇到了类似的问题, 因此他们的底层存储 [tikv](https://github.com/pingcap/tikv) 也选择参考 `etcd/raft` 的实现.



我们先简单介绍一下 `etcd/raft` .



##### raftpb

[raftpb](https://github.com/coreos/etcd/tree/master/raft/raftpb) 使用 protobuf 定义了基础数据结构.



##### Node

这个接口定义基本说明了使用者能做哪些事情...

```
// Node represents a node in a raft cluster.
type Node interface {
	// Tick increments the internal logical clock for the Node by a single tick. Election
	// timeouts and heartbeat timeouts are in units of ticks.
	Tick()
	// Campaign causes the Node to transition to candidate state and start campaigning to become leader.
	Campaign(ctx context.Context) error
	// Propose proposes that data be appended to the log.
	Propose(ctx context.Context, data []byte) error
	// ProposeConfChange proposes config change.
	// At most one ConfChange can be in the process of going through consensus.
	// Application needs to call ApplyConfChange when applying EntryConfChange type entry.
	ProposeConfChange(ctx context.Context, cc pb.ConfChange) error
	// Step advances the state machine using the given message. ctx.Err() will be returned, if any.
	Step(ctx context.Context, msg pb.Message) error

	// Ready returns a channel that returns the current point-in-time state.
	// Users of the Node must call Advance after retrieving the state returned by Ready.
	//
	// NOTE: No committed entries from the next Ready may be applied until all committed entries
	// and snapshots from the previous one have finished.
	Ready() <-chan Ready

	// Advance notifies the Node that the application has saved progress up to the last Ready.
	// It prepares the node to return the next available Ready.
	//
	// The application should generally call Advance after it applies the entries in last Ready.
	//
	// However, as an optimization, the application may call Advance while it is applying the
	// commands. For example. when the last Ready contains a snapshot, the application might take
	// a long time to apply the snapshot data. To continue receiving Ready without blocking raft
	// progress, it can call Advance before finishing applying the last ready.
	Advance()
	// ApplyConfChange applies config change to the local node.
	// Returns an opaque ConfState protobuf which must be recorded
	// in snapshots. Will never return nil; it returns a pointer only
	// to match MemoryStorage.Compact.
	ApplyConfChange(cc pb.ConfChange) *pb.ConfState

	// TransferLeadership attempts to transfer leadership to the given transferee.
	TransferLeadership(ctx context.Context, lead, transferee uint64)

	// ReadIndex request a read state. The read state will be set in the ready.
	// Read state has a read index. Once the application advances further than the read
	// index, any linearizable read requests issued before the read request can be
	// processed safely. The read state will have the same rctx attached.
	ReadIndex(ctx context.Context, rctx []byte) error

	// Status returns the current status of the raft state machine.
	Status() Status
	// ReportUnreachable reports the given node is not reachable for the last send.
	ReportUnreachable(id uint64)
	// ReportSnapshot reports the status of the sent snapshot.
	ReportSnapshot(id uint64, status SnapshotStatus)
	// Stop performs any necessary termination of the Node.
	Stop()
}
```



##### Storage

这是日志持久化层

但是实际上可以看到, 没有要求提供写的方法.

言下之意是 *我只需要读, 至于该怎么存, 存哪里, 请在 Node.Ready() 中自行解决*

```
// Storage is an interface that may be implemented by the application
// to retrieve log entries from storage.
//
// If any Storage method returns an error, the raft instance will
// become inoperable and refuse to participate in elections; the
// application is responsible for cleanup and recovery in this case.
type Storage interface {
	// InitialState returns the saved HardState and ConfState information.
	InitialState() (pb.HardState, pb.ConfState, error)
	// Entries returns a slice of log entries in the range [lo,hi).
	// MaxSize limits the total size of the log entries returned, but
	// Entries returns at least one entry if any.
	Entries(lo, hi, maxSize uint64) ([]pb.Entry, error)
	// Term returns the term of entry i, which must be in the range
	// [FirstIndex()-1, LastIndex()]. The term of the entry before
	// FirstIndex is retained for matching purposes even though the
	// rest of that entry may not be available.
	Term(i uint64) (uint64, error)
	// LastIndex returns the index of the last entry in the log.
	LastIndex() (uint64, error)
	// FirstIndex returns the index of the first log entry that is
	// possibly available via Entries (older entries have been incorporated
	// into the latest Snapshot; if storage only contains the dummy entry the
	// first log entry is not available).
	FirstIndex() (uint64, error)
	// Snapshot returns the most recent snapshot.
	// If snapshot is temporarily unavailable, it should return ErrSnapshotTemporarilyUnavailable,
	// so raft state machine could know that Storage needs some time to prepare
	// snapshot and call Snapshot later.
	Snapshot() (pb.Snapshot, error)
}
```



##### 网络传输

`etcd/raft` 没有提供任何网络传输层的接口定义.

与日志的持久化类似, *我只告诉你哪些 message 需要发出, 怎么发, 发往哪里请自行解决* 😂.



##### 总得有一些开箱即用的东西...

对于日志持久化层, `etcd/raft` 提供了一个内存版本的 `Storage` 实现 `MemoryStorage` , 通过 [wal](https://github.com/coreos/etcd/tree/master/wal) 落盘.

而 [rafthttp](https://github.com/coreos/etcd/tree/master/rafthttp) 则提供了节点寻址和基于 http 的网络传输能力...



可以参考一下 etcd 官方提供的 [demo](https://github.com/coreos/etcd/tree/master/contrib/raftexample) .

港真, 我在用 `hashicorp/raft` 写了一些基本能用的小玩具之后看这个 demo, 还是把我绕晕了.



#### 0.3 dgraph

[dgraph](https://github.com/dgraph-io/dgraph) 是一款使用 go 语言开发的分布式图数据库.



##### dgraph zero

zero 节点用于管理 dgraph 集群, 维护成员信息, 数据的 sharding 和 rebalancing.

我们借着阅读 zero 的实现代码来看一看 `etcd/raft` 的使用, 以及它的周边组件的实现方式.

