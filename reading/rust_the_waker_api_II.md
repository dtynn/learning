## The Waker API II: waking across threads

原文: [The Waker API II: waking across threads](https://boats.gitlab.io/blog/post/wakers-ii/)



- 讲述了设计出 Waker 和 LocalWaker 两种类型的初衷.
- 对二者进行区分带来的认知成本
- 结论:
  1. 针对相同线程的优化, 可以藉由 TLS 来解决
  2. 对于仅支持从同一个线程唤醒的场景, 非原子性的引用计数 Rc 无法胜任, 可以通过:
     - 使用 TaskID 的实现方式
     - 使用原子性引用计数