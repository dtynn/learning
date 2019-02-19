

## The Waker API I: what does a waker do?

原文: [The Waker API I: what does a waker do?](https://boats.gitlab.io/blog/post/wakers-i/)



### The poll/wait/wake cycle

正确使用 async / await 特性的程序包含三个基本组件:

- futures: 由终端用户编写的可停顿的代码块

- executor: 用于 futures 的调度

- event sources: 在合适的时机唤醒 executor

  在异步 IO 中, IO 事件的 e-s 通常被称为 reactor



futures 的执行阶段:

1. Poll: 执行到某个中断点
2. Wait: 等待某个事件
3. Wake: 事件发生, 恢复执行



一种抽象的理解方式:

- executor 管理程序的计算资源
- reactor 管理程序的 IO 资源
- executor 和 reactor 构成 event loop



rust 只提供 API 的定义, 使用者可以自行组合不同的 executor 和 reactor



### Requirements on the Waker API

用于 executor 和 reactor 共同协调 waiting 和 waking 发生的时机.



#### What does the “poll” phase require?

- **dynamically dispatched**



#### What does the “wait” phase require?

- **the waker type must implement Clone**

  原因: event source 注册时需要保存 waker, 而 waker 需要能够同时处理不同的事件, 因此不能被某一个 event resource 独立拥有.



#### What does the “wake” phase require?



### Waking and threading



### Conclusion

