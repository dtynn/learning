### core::nodes

`nodes` 中主要是底层的网络原语。



#### core::nodes::collection

`collection` 对网络端点间的通信进行了抽象定义， 主要涵盖几个方面：

1. 端点管理
2. 连接管理
3. 事件管理



各个定义的概述如下：

- `CollectionStream` 是 core::nodes::collection 中的核心结构体，其定义如下：

  ```rust
  pub struct CollectionStream<TInEvent, TOutEvent, THandler, TReachErr, THandlerErr, TUserData, TConnInfo = PeerId, TPeerId = PeerId> {
  
      inner: tasks::Manager<TInEvent, TOutEvent, THandler, TReachErr, THandlerErr, TaskState<TConnInfo, TUserData>, TConnInfo>,
  
      nodes: FnvHashMap<TPeerId, TaskId>,
  }
  
  
  ```

  从结构体定义我们大致可以推断出 `CollectionStream` 的用途是管理端点，并处理端点之间的通信事件 `CollectionEvent`。

- 本地端点可以向远端发起连接尝试，会返回一个 `ReachAttemptId`；

  必要的时候也可以通过 `ReachAttemptId` 来中断一次连接尝试，根据中断的结果， 会产生一个 `InterruptedReachAttempt` 或 `InterruptError`；

- `CollectionReachEvent` 是远端向本端发起的连接请求，本端可以接受 (accept) 或 拒绝 (deny)；

  一旦接受，会转化成一个 `CollectionNodeAccept`；

- 连接成功的两端会分别将对端转化成己方的一个 `PeerMut`；

  对于已经连接的 `PeerMut` 可以执行通信或中断操作；



#### core::nodes::listeners

`listeners` 中的定义比较简单：

- `ListenersEvent`， 监听事件， 包含：

  - `NewAddress`：新的监听地址；
  - `AddressExpired`：过期的监听地址；
  - `Incoming`：连接；
  - `Closed`：关闭的 Listener

- `ListenersStream`：

  ```rust
  pub struct ListenersStream<TTrans>
  where
      TTrans: Transport,
  {
      transport: TTrans,
  
      listeners: VecDeque<Listener<TTrans>>
  }
  
  
  struct Listener<TTrans>
  where
      TTrans: Transport,
  {
  
      listener: TTrans::Listener,
  
      addresses: SmallVec<[Multiaddr; 4]>
  }
  
  ```

   `ListenersStream` 可以看做是一组 使用同一种 `Transport` 的，监听不同地址的 `Listener` 的集合，它会持续产生 `ListenerEvent` 流。



#### core::nodes::node

`node` 中的核心定义也是两个：

- `NodeEvent`:

  - `InboundSubstream`
  - `OutboundSubStream`

- `NodeStream`：`NodeEvent` 的流

  ```rust
  pub struct NodeStream<TMuxer, TUserData>
  where
      TMuxer: muxing::StreamMuxer,
  {
  
      muxer: Arc<TMuxer>,
  
      outbound_substreams: SmallVec<[(TUserData, TMuxer::OutboundSubstream); 8]>,
  }
  
  ```

  

可以简单认为：

1. 在协议支持的情况下， 可以通过 `StreamMuxer` 在一个连接上维护多个相互独立的子数据流 (`SubStream`)；
2. `NodeStream` 管理所有子流上的数据发送和接收；



#### core::nodes::handled_node

`handled_node` 是对 `node` 的一层封装。

可以认为 `node` 中传递的是原始数据， 而 `handled_node` 则承担更高层逻辑中的事件 `Event` 和原始数据之间的转换工作。

`handled_node` 中主要包含两块定义：

1. `NodeHandlerEvent` 和 `NodeHandler`， 即 `handled_node` 中的 `handled` 部分；
   - `NodeHandlerEvent` 包含两类定义：
     1. `OutboundSubstreamRequest` 向指定目标发起一个子数据流的事件；
     2. `Custom` 其他类型的事；
   - `NodeHandler` 是一个 trait，具备如下的能力：
     1. 处理子数据流；
     2. 处理来自外界的事件(Self::InEvent)；
     3. 持续产生 `NodeHandlerEvent` (Self::OutEvent) 流；
2. `HandledNode` 是一个 `NodeStream` 和一个 `NodeHandler` 的组合，主要职责为：
   1. 通过 `NodeHandler` 处理来自外界的事件；
   2. 将 `NodeHandler` 中产生的 `NodeHandlerEvent::OutboundSubstreamRequest` 转发给 `NodeStream` 形成子数据流的实际创建动作；
   3. 将 `NodeHandler` 中产生的 `NodeHandlerEvent::Custom` 转化成持续的事件流；



#### core::nodes::tasks

`tasks` 也包含两块核心定义：`Task` 本身， 和管理 `Task` 的 `Manager`。



##### Task

 `Task` 是定义在模块内部的。

```rust
/// Message to transmit from the public API to a task.
pub enum ToTaskMessage<T> {
    /// An event to transmit to the node handler.
    HandlerEvent(T),
    /// When received, stores the parameter inside the task and keeps it alive
    /// until we have an acknowledgment that the remote has accepted our handshake.
    TakeOver(mpsc::Sender<ToTaskMessage<T>>)
}

/// Message to transmit from a task to the public API.
pub enum FromTaskMessage<T, H, E, HE, C> {
    /// A connection to a node has succeeded.
    NodeReached(C),
    /// The task closed.
    TaskClosed(Error<E, HE>, Option<H>),
    /// An event from the node.
    NodeEvent(T)
}

pub struct Task<F, M, H, I, O, E, C>
where
    M: StreamMuxer,
    H: IntoNodeHandler<C>,
    H::Handler: NodeHandler<Substream = Substream<M>>
{
    /// The ID of this task.
    id: TaskId,

    /// Sender to transmit messages to the outside.
    sender: mpsc::Sender<(FromTaskMessage<O, H, E, <H::Handler as NodeHandler>::Error, C>, TaskId)>,

    /// Receiver of messages from the outsize.
    receiver: stream::Fuse<mpsc::Receiver<ToTaskMessage<I>>>,

    /// Inner state of this `Task`.
    state: State<F, M, H, I, O, E, C>,

    /// Channels to keep alive for as long as we don't have an acknowledgment from the remote.
    taken_over: SmallVec<[mpsc::Sender<ToTaskMessage<I>>; 1]>
}

enum State<F, M, H, I, O, E, C>
where
    M: StreamMuxer,
    H: IntoNodeHandler<C>,
    H::Handler: NodeHandler<Substream = Substream<M>>
{
    /// Future to resolve to connect to the node.
    Future {
        /// The future that will attempt to reach the node.
        future: F,
        /// The handler that will be used to build the `HandledNode`.
        handler: H,
        /// While we are dialing the future, we need to buffer the events received on
        /// `receiver` so that they get delivered once dialing succeeds. We can't simply leave
        /// events in `receiver` because we have to detect if it gets closed.
        events_buffer: Vec<I>
    },

    /// An event should be sent to the outside world.
    SendEvent {
        /// The node, if available.
        node: Option<HandledNode<M, H::Handler>>,
        /// The actual event message to send.
        event: FromTaskMessage<O, H, E, <H::Handler as NodeHandler>::Error, C>
    },

    /// We started sending an event, now drive the sending to completion.
    ///
    /// The `bool` parameter determines if we transition to `State::Node`
    /// afterwards or to `State::Closing` (assuming we have `Some` node,
    /// otherwise the task will end).
    PollComplete(Option<HandledNode<M, H::Handler>>, bool),

    /// Fully functional node.
    Node(HandledNode<M, H::Handler>),

    /// Node closing.
    Closing(Close<M>),

    /// Interim state that can only be observed externally if the future
    /// resolved to a value previously.
    Undefined
}



```



通过 `Task` 的 `Future` 实现, 我们可以了解 `State` 的细节:

- `State::Future` :

  1. 持续接收 `ToTaskMessage::HandleEvent` 和 `ToTaskMessage::TakeOver` 并缓存下来, 直到 receiver 中暂时没有更多信息.
  2. 确认与目标 node 的连接情况:

  - 连接成功
    1. 使用连接信息 (连接信息 conn_info 和 连接复用器 muxer) 构造成 `HandledNode`
    2. 将 `events_buffer`  中的 event 全部转发到 node 中
    3. 将状态转换成  `State::SendEvent`, 附带 event 为 `FromTaskMessage::NodeReached`
  - 未连接: 返回 `Async::NotReady`, 等待下一次 poll
  - 连接失败: 将状态转换成 `State::SendEvent`, 附带 event 为 `FromTaskMessage::TaskClosed`

- `State::SendEvent` :

  1. 持续接收 `ToTaskMessage`:

     - 若接收到 `ToTaskMessage::TakeOver`, 则保存
     - 若接收到 `ToTaskMessage::HandleEvent` 且存在已连接的 node, 则向 node 发送

     - 若 receiver 中暂时没有更多信息, 则中断接收
     - 若 receiver 关闭, 根据是否存在已连接的 node, 将状态转为 `State::Closing` 或结束当前任务

  2. 通过 sender 将附加的 event 向外界发送, 将 event 插入 sender 队列:

     - 入队成功: 将状态转换成 `State::PollComplete`

     - 未入队: 维持状态等待下一次 poll
     - 入队异常: 根据是否存在已连接的 node, 将状态转为 `State::Closing` 或结束当前任务

- `State::Closing` : 维持状态, 持续等待直到 `State::Future` 中连接成功获取的 muxer 关闭

- `State::PollComplete` :

  1. 持续接收 `ToTaskMessage`, 处理逻辑同 `State::SendEvent`;
  2. 确认 sender 中 event 的发送状态:
     - 未发送: 维持状态等待下一次 poll
     - 发送成功: 根据 event 类型决定下一个状态:
       - `FromTaskMessage::TaskClosed`: 状态转换成 `State::Closing`
       - 其他: 状态转换成 `State::Node`
     - 发送异常: 根据是否存在已连接的 node, 将状态转为 `State::Closing` 或结束当前任务

- `State::Node`: 

  在此状态下, 已经确保存在已连接的 node, 

  1. 持续接收 `ToTaskMessage`, 处理逻辑同 `State::SendEvent`
  2. 若已经接收到 node 的确认信息, 释放所有 `taken_over` 中的发送端, 通知接收端
  3. 尝试获取 node 传递过来的 event:
     - 未就绪: 维持状态等待下一次 poll
     - 就绪: 将状态转换成 `State::SendEvent`, 附带 event 为 `FromTaskMessage::NodeEvent`
     - 异常: 将状态转换成 `State::SendEvent`, 附带 event 为 `FromTaskMessage::TaskClosed`



综上, 我们可以总结 `Task` 的工作方式如下:

1. 尝试连接指定 node, 根据连接情况, 向外界发送 event
2. 在 node 连接成功的情况下, 持续在外界和 node 之间传递 event
3. 当外界停止发送信息, 或通信过程中出现异常时, 进入清理并关闭当前 task 的流程



##### Manager

`Manager`的作用是:

1. 管理并驱动所有 `Task`
2. 持续分发外界流入 和 `Task` 反馈的 event





#### core::nodes::network

`network` 将 `core::nodes` 中的其他模块和概念进行整合, 构成完整的 p2p 网络底层框架.



`Network` 是核心结构体, 其定义如下:

```rust
/// Implementation of `Stream` that handles the nodes.
pub struct Network<TTrans, TInEvent, TOutEvent, THandler, THandlerErr, TConnInfo = PeerId, TPeerId = PeerId>
where
    TTrans: Transport,
{
    /// Listeners for incoming connections.
    listeners: ListenersStream<TTrans>,

    /// The nodes currently active.
    active_nodes: CollectionStream<TInEvent, TOutEvent, THandler, InternalReachErr<TTrans::Error, TConnInfo>, THandlerErr, (), (TConnInfo, ConnectedPoint), TPeerId>,

    /// The reach attempts of the network.
    /// This needs to be a separate struct in order to handle multiple mutable borrows issues.
    reach_attempts: ReachAttempts<TPeerId>,

    /// Max numer of incoming connections.
    incoming_limit: Option<u32>,

    /// Unfinished take over message to be delivered.
    ///
    /// If the pair's second element is `AsyncSink::NotReady`, the take over
    /// message has yet to be sent using `PeerMut::start_take_over`.
    ///
    /// If the pair's second element is `AsyncSink::Ready`, the take over
    /// message has been sent and needs to be flushed using
    /// `PeerMut::complete_take_over`.
    take_over_to_complete: Option<(TPeerId, AsyncSink<InterruptedReachAttempt<TInEvent, (TConnInfo, ConnectedPoint), ()>>)>
}

```



`Network` 能够

1. 接收来自其他节点的连接请求(listeners)
2. 尝试连接指定的节点 (active_nodes)
3. 将 1, 2 中创建的连接转换成本端与对端的信息通道 (active_nodes), 并加以维护



这样外界, 如上层应用, 即可通过  `Network`  与其他端通信:

1. 由  `Network` 负责信息的传递
2. 由应用层负责信息的解读和执行

