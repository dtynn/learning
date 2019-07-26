### mod core

`core` 模块中包含了 `libp2p` 中主要类型的定义。



#### core::either

 `either` 中定义了一系列用于选择、切换、fallback 场景的枚举量。



#### core::identity

 `identity` 主要提供了用于身份鉴定的密钥生成方法。



#### core::multiaddr

`multiaddr` 主要提供了 `MultiAddr` 类型的定义、生成和解析方法，以及错误信息的定义。



#### core::muxing

`muxing` 主要定义了 trait `StremMuxer`。

`StreamMuxer` 用于对一个连接进行多路复用， 可以将一个连接分成多个相互独立互不干扰的子数据流。

某些协议可以通过实现 `StreamMuxer` 达到以极低的资源开销达到承载大量请求的目的，类似 `http2` 在一个 tcp 连接上处理多个 http 请求。



完整的`StremMuxer` 定义如下：

```
pub trait StreamMuxer {
    
    type Substream;

    type OutboundSubstream;

    type Error: Into<io::Error>;

    fn poll_inbound(&self) -> Poll<Self::Substream, Self::Error>;

    fn open_outbound(&self) -> Self::OutboundSubstream;

    fn poll_outbound(&self, s: &mut Self::OutboundSubstream) -> Poll<Self::Substream, Self::Error>;

    fn destroy_outbound(&self, s: Self::OutboundSubstream);

    fn read_substream(&self, s: &mut Self::Substream, buf: &mut [u8]) -> Poll<usize, Self::Error>;

    unsafe fn prepare_uninitialized_buffer(&self, buf: &mut [u8]) -> bool {
        for b in buf.iter_mut() { *b = 0; }
        true
    }

    fn write_substream(&self, s: &mut Self::Substream, buf: &[u8]) -> Poll<usize, Self::Error>;

    fn flush_substream(&self, s: &mut Self::Substream) -> Poll<(), Self::Error>;

    fn shutdown_substream(&self, s: &mut Self::Substream) -> Poll<(), Self::Error>;

    fn destroy_substream(&self, s: Self::Substream);

    fn is_remote_acknowledged(&self) -> bool;

    fn close(&self) -> Poll<(), Self::Error>;

    fn flush_all(&self) -> Poll<(), Self::Error>;
}


```

我们可以简单地认为 `StreamMuxer::Substream` 是来自远端的请求， 而 `StreamMuxer::OutboundStream` 则是本地向远端发起的请求。



#### core::transport

`transport` 中主要提供的就是 trait `Transport` 了，在 major concepts 里已经做了解读。



#### core::nodes & core::upgrade

`nodes` 和 `upgrade` 将单独解读。

