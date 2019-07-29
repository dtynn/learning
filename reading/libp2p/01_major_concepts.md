### Marjo concepts

#### Multiaddr

`multiaddr` 包含地址信息和所使用的协议栈，用于节点之间创建连接。



#### Transport & Connection Upgrades

##### 定义

- Transport

  `Transport` 是一个 Trait， 代表面向连接的通信通道。它要求至少提供基于 multiaddr 的监听和拨号实现。

- Connection Upgrades

  `Connection Upgrades` 是将已有的`Transport` 转换成一个遵循某个附加协议的新的 `Transport` 的过程。

  例如 web 开发中常见的 websocket， 就可以理解为由一个 tcp 连接经 upgrades 而得。



##### 代码

- 类型定义：

  trait `Transport` 包含了五种关联类型：

  ```rust
  pub trait Transport {
  
      type Output;
      
      type Error: error::Error;
      
      type Listener: Stream<Item = ListenerEvent<Self::ListenerUpgrade>, Error = Self::Error>;
      
      type ListenerUpgrade: Future<Item = Self::Output, Error = Self::Error>;
      
      type Dial: Future<Item = Self::Output, Error = Self::Error>;
  
      ......
  }
  ```

  其中， `Listener` 还涉及了另外一个枚举量 `ListenerEvent`

  ```rust
  #[derive(Clone, Debug, PartialEq)]
  pub enum ListenerEvent<T> {
      /// The transport is listening on a new additional [`Multiaddr`].
      NewAddress(Multiaddr),
      /// An upgrade, consisting of the upgrade future, the listener address and the remote address.
      Upgrade {
          /// The upgrade.
          upgrade: T,
          /// The listening address which produced this upgrade.
          listen_addr: Multiaddr,
          /// The remote address which produced this upgrade.
          remote_addr: Multiaddr
      },
      /// A [`Multiaddr`] is no longer used for listening.
      AddressExpired(Multiaddr)
  }
  
  
  ```

  从类型定义和约束，大致可以看出：

  - `Output` 代表一个具体的连接类型；
  - `Error` 是连接创建过程中产生的异常；
  - `Dial` 是一个从本地向远端拨号的过程，最终会创建一个 `Output` 类型的连接；
  - `Listener` 用于本地监听， 会生成一系列的`ListenerEvent` 事件；
  - `ListenerUpgrade` 是由 `ListenerEvent::Upgrade` 类型的事件转换成一个从远端向本地拨号发起的`Output`连接的过程；

  我们可以类比 tcp 连接的相关概念来进行理解。

- 方法定义：

  `Transport` 定义中有一些比较值得关注的方法：

  ```rust
  pub trait Transport {
      ......
  
      fn listen_on(self, addr: Multiaddr) -> Result<Self::Listener, TransportError<Self::Error>>
      where
          Self: Sized;
  
      fn dial(self, addr: Multiaddr) -> Result<Self::Dial, TransportError<Self::Error>>
      where
          Self: Sized;
  
      ......
  
      fn with_upgrade<U, O, E>(self, upgrade: U) -> Upgrade<Self, U>
      where
          Self: Sized,
          Self::Output: AsyncRead + AsyncWrite,
          U: InboundUpgrade<Self::Output, Output = O, Error = E>,
          U: OutboundUpgrade<Self::Output, Output = O, Error = E>
      {
          Upgrade::new(self, upgrade)
      }
      
      ......
  }
  
  ```

  - `listen_on` 用于创建 `Self::Listener` 对象，`dial` 用于创建 `Self::Dial` 对象；
  - `with_upgrade` 对应 `Connection Upgrades` 概念， 用于对 `Self::Output` 类型的连接进行基于指定协议的封装， 根据类型约束来看， 流入数据和流出数据可以使用不同的协议；



由此我们可以看出，trait  `Transport` 主要关注：

- 接受来自远端的连接；
- 向远端发起连接；
- 处理数据协议；



### Network Behavior & Swarm

- `Network Behavior` 定义了一个节点在整个网络中可以执行的某些动作，比如：
  - 周期性 ping 其他节点；
  - 周期性获取其他节点的信息；
  - 请求 DHT 信息并向其他节点传播；
- `Swarm` 管理着本节点和其他节点之间的连接，以及与其他节点之间的通信行为。





