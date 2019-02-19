## Generators I: Toward a minimum viable product

[原文](https://boats.gitlab.io/blog/post/generators-i/)



### Defining the boundaries of our use case

两类 generators 的基本使用场景:

- 实现 "coroutine"
- 表达某种可迭代的资源

第二种场景是当前主要的目标



### Dividing up the design space

明确主要目标之后, 可以把当前没解决的问题分成两块:

- 为了使 generator 功能稳定而产生的问题, 当前
- 针对我们的场景, 不需要立即解决, 稍后

#### 当前:

- 语法
- 相关 trait 和 type 的签名
- 与Iterator 和 Stream 类型的相互转化
- **Self-referential and Unpin generators**



#### 稍后

- 返回 非 () 类型的 generator
- 接收 resumption 参数的 generator

