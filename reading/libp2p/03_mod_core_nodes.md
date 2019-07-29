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

 `Task` 是定义在模块内部的，我们主要看他的代码。

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







#### core::nodes::network



