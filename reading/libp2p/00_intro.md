### Intro

#### libp2p

[libp2p](https://libp2p.io/) 是一个模块化的 p2p 网络栈库。

官方宣称它的包含以下特性：

- User several transport
- Native roaming
- Runtime freedom
- Protocol muxing
- Work offline
- Encrypted connections
- Upgrade without compromises
- Work in the browser
- Good for high latency scenarios

整体来看， 关注点集中在安全性， 灵活性， 可靠性等方面。



#### rust-libp2p

[rust-libp2p](https://github.com/libp2p/rust-libp2p) 是 libp2p 的 rust 版实现， 目前仍处于 [WIP 状态](https://github.com/libp2p/rust-libp2p/graphs/commit-activity)。

这里我将以 rust-libp2p v0.11.0 的[文档]((https://docs.rs/libp2p/0.11.0/libp2p/)) 为参考， 通过阅读代码的方式来尝试理解 libp2p。

代码的提交版本为 bcc7c4d349。