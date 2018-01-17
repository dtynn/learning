#### Summary

`prometheus/tsdb`  (下称 ptsdb ) 的结构体之家你的层次大概可以这样划分:

- DB: 对外提供的核心对象

  - Block 已经持久化的, 覆盖某个时间段的时序数据. Block 的
    - Index: 用于保存 labels 的索引数据
    - Chunk: 用于保存时间戳-采样值 数据


  - Head: 由于 ptsdb 规定, 数据必须增序写入, 已经持久化的 Block 不能再写入, 因此一个时刻只会有一个可供写入的 Block, 即 Head. Head 同时还承担记录删除动作的任务
    - WAL 增删改的动作都会先进入 WAL, 供后续恢复用
    - Tombstone: 用于标记删除动作, 被标记的数据在 compact 的时候统一清理
  - Compactor: 对文件进行压缩. Block 数据的组织参考了 LSM, 因此 Compactor 的实现也和基于 LSM 的 kv db 类似.



关于 ptsdb, [时间序列数据的存储和计算 - 开源时序数据库解析（四）](https://zhuanlan.zhihu.com/p/32900004) 这篇文章有更宏观的阐述, 可以参考.



