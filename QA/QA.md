#### goroutine 的调度

M-P-G

M: OS 线程, 称之为 Machine, 是实际的执行体, 不保持状态

G: goroutine, 仅仅保存任务状态, 包括栈, 指令指针, 以及其他信息, 比如阻塞它的 channel

P: processor



P 有一个 runqueue, 用于存放准备好被调度的 G, 全局还有一个 runqueue

P 不断循环取出 runqueue 中的 G, 交给 M 执行.

M  拿到 G 之后, 取出 G 中保存的状态, 继续执行. 如果需要中途切换, 则将相关寄存器值保存回 G. 任何 M 都可以恢复执行这个 G.



P 可以重新绑定 M



#### 抢占式调度

协作式多任务处理: 下一个进程被调度的前提是当前进程主动放弃

抢占式多任务处理: 操作系统可以剥夺耗时长的时间片, 提供给其他进程.



#### TCP/UDP 区别

链接 vs 无连接
流 vs 报文
数据可靠 vs 不可靠
有序 vs 无序



#### 堆和栈

数据结构

- 堆 是优先队列,
- 栈是后进先出



程序运行时

- 栈: 随程序调用分配, 存放程序的执行状态和部分数据
- 堆: 在进程的整个未分配空间中分配内存, 存放栈中放不下的数据, 以及生命周期不确定的数据



#### 程序在跨线程、跨进程、跨主机间的通信方式分别有哪几种方式？

跨线程:

- 共享变量
- 管道

跨进程

- 管道
- 信号
- 消息队里
- 共享内存
- 信号量
- 套接字

跨机器

- 套接字



#### 怎样判断一个函数是线程安全的

看函数内的资源是否会被多个线程同时使用, 尤其是写操作



#### 网络字节序和主机字节序分别是什么？

主机字节序: 整数在内存中保存的顺序, LE BE 都有

网络字节序: TCP/IP 中规定好的数据表示格式, BE



#### base64, md5,  rsa有什么区别

base64 编码格式

md5 摘要算法

rsa 非对称加密, 公钥加密, 私钥解密



#### https 通信流程

1. 客户端请求, 服务端返回公钥和证书
2. 客户端验证公钥
3. 客户端生成随机的对称加密密钥, 通过公钥加密后发送给服务端
4. 服务端解密得到加密密钥
5. 通过加密密钥通信



#### 聚簇索引

聚簇索引的叶子节点是数据节点, 通常只有主键使用



#### 关系型数据库范式

- 第一范式: 所有字段都是不可再分的
- 第二范式: 字段依赖于全部属性
- 第三范式: 字段互不依赖
- BC 范式: 主属性内部不能部分或传递依赖
- 第四范式: 消除表中的多值依赖

多值依赖: 一个关系, 至少存在A, B, C 三种属性,  对于每一个 A 值, 有一组确定的 B 和 C, 且 B 独立于 C



#### MongoDB 索引类型和属性

单键, 复合, 多键, Geo, 全文, Hash



索引属性: TTL, 唯一, 稀疏



#### SOLID Go Design

##### Single Responsibility Principle 单一功能原则

A class should have one, and only one, reason to change.

鼓励你在package中构建functions、types以及方法表现出自然的凝聚力。



##### Open / Closed Principle 开闭原则

Software entities should be open for extension, but closed for modification. 

鼓励你使用嵌入将简单的type组合成更为复杂的。



##### Liskov Substitution Principle 里氏替换原则

子类能够替换其超类被使用

鼓励你在package之间表达依赖关系时用interface，而非具体类型。通过定义小巧的interface，我们可以更有信心地切实满足其合约。



##### Interface Segregation Principle 接口隔离原则

Clients should not be forced to depend on methods they do not use.

鼓励你仅取决于所需行为来定义函数和方法。如果你的函数仅仅需要有一个方法的interface做为参数，那么它很有可能只有一个责任。



##### Dependency Inversion Principle 依赖反转原则

High-level modules should not depend on low-level modules. Both should depend on abstractions. Abstractions should not depend on details. Details should depend on abstractions.

鼓励你在编译时将package所依赖的东西移除



#### golang 在遍历 slice, map 过程中插入/删除元素

##### slice

遍历长度只在循环外获取一次, 该长度决定了遍历次数

- 增加元素, 会改变 slice 内的元素, 但不会改变遍历次数
- 删除元素, 会改变 slice 内的元素, 但不会改变遍历次数



##### map

- 删除元素, 则对应的元素不再出现在迭代过程中
- 插入元素, 可能出现, 也可能不出现, 不可控



#### 关于 defer

- 执行顺序 FILO
- 在 defer 出现的行调用, 在 return 后执行 -> 参数值
- 对返回值的影响, 遵循值传递, 指针传递
- 对返回值的影响: 如果是返回值定义了变量名, defer 内赋值, 可能影响