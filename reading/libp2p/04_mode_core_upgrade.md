### core::upgrade

`upgrade` 主要包含对连接 (connection) 或 子流(substream) 进行协议封装和转换所需的类型和定义.

#### Trait UpgradeInfo

用于告知外界自己能够支持的协议(名称)列表

```rust
pub trait ProtocolName {

    fn protocol_name(&self) -> &[u8];
}

pub trait UpgradeInfo {

    type Info: ProtocolName + Clone;

    type InfoIter: IntoIterator<Item = Self::Info>;

    fn protocol_info(&self) -> Self::InfoIter;
}


```





####  Trait Upgrade

流入和流出分开进行了定义:

- InboundUpgrade

  ```rust
  pub trait InboundUpgrade<C>: UpgradeInfo {
  
      type Output;
  
      type Error;
  
      type Future: Future<Item = Self::Output, Error = Self::Error>;
  
      fn upgrade_inbound(self, socket: Negotiated<C>, info: Self::Info) -> Self::Future;
  }
  
  
  ```

  

- OutboundUpgrade

  ```rust
  pub trait OutboundUpgrade<C>: UpgradeInfo {
  
      type Output;
  
      type Error;
  
      type Future: Future<Item = Self::Output, Error = Self::Error>;
  
      fn upgrade_outbound(self, socket: Negotiated<C>, info: Self::Info) -> Self::Future;
  }
  
  ```

- InboundUpgradeExt & OutboundUpgradeExt:

  在 InboundUpgrade 和 OutboundUprade 的基础上, 增加了对 `Output` 和 `Error` 进行 map 的方法定义.



可以看到, 核心的类型都是 `Negotiated<C>`,  其中:

- `C` 具有类型约束 `Read + AsyncRead + Write + AsyncWrite`
- `Negotiated` 是对 `C` 的一层封装, 同样实现了 `Read + AysncRead + Write + AsyncWrite`



从命名和约束上我们大致可以推断, `InboundUpgrade` 和 `OutboundUpgrade` 是对已经完成了协议选择和握手确认的读写通道进行指定封装的数据类型.



#### Structs

主要定义了三类 `Upgrade`:

1. `DeniedUpgrade`
2. `OptionalUpgrade`
3. `SelectUpgrade`



它们都同时实现了 `UpgradeInfo`, `InboundUpgrade` 和 `OutboundUpgrade`

