在目前的 db-engines 时序数据库排名中, 第一位的 [influxdb](https://github.com/influxdata/influxdb) 和第七位的 [prometheus](https://github.com/prometheus/prometheus) 都是由 go 实现的.

这次我打算阅读的源码是 promethus 的组件 [prometheus/tsdb](https://github.com/prometheus/tsdb)



##### 时序数据库

首先简单介绍一下, 时序数据库 (tsdb) 全称为 Time Series Database, 通常作为各类监测, 监控系统的核心组成部分. 通常需要极高的写入速度, 尽量经济的空间占用, 和尽可能高的批量数据查询效率.

除了上面提到的 influxdb 和 prometheus, 比较典型的还有 OpenTSDB,  KairosDB 等.



##### 时序数据库的存储

各类时序数据库的数据结构基本都以 tags , timestamp 和 value 这样一组概念为核心, 称为数据点, 类似一般数据库中的记录或条目概念.

查询大多基于时间戳, 并以 tags 作为辅助筛选项.

至于底层存储, 则有不同的流派, 一类是使用已有的成熟数据库产品, 比如 OpenTSDB 默认使用 HBase, KairosDB 基于 Cassandra 等; 另一类则采用自研的存储引擎, 比如 Influxdb 在  LSM 的基础上针对时序数据的特性开发了 TSM 存储引擎 tsm1, 而 prometheus/tsdb 则参考了 facebook 的 论文 [Gorilla: A Fast, Scalable, In-Memory Time Series Database](http://www.vldb.org/pvldb/vol8/p1816-teller.pdf)



##### Facebook 的 Gorilla

Gorilla 的 最大特点, 是快和高压缩率.

facebook 总结了自身海量的 IT 系统监控数据得出以下结论

1. 相邻数据点时间戳的差值相对固定, 即使变化也仅在一个很小的范围内浮动;
2. 相邻数据点的值变化幅度也很小, 且有相当比例的值变化为0;
3. 热点数据查询的频率远远超出非热点数据, 且越是近期的数据热度越高.



继而提出了 Gorilla 的存储设计:

- 将一段时间的数据点存储在一个数据块内;
- 数据块仅第一个点记录时间戳和值的绝对值, 之后所有点仅记录其值的变化情况.



以传统时序数据结构中, 数据点的时间戳和值均记为 64 bit 的值, 共占用 16 Byte 为参照, Gorilla 疯狂地将每数据点的平均空间压缩到1.37 Byte, 与此同时满足了 fb 自身 700M 每分钟的数据点写入需求.



Gorilla 的缺点则在查询上, 对于目标时间点, 需要将其所在的数据块整个取出, 进行遍历.

因此, facebook 将其热点数据完全放在内存之中.



facebook 开源了 Gorilla 的 C++ 实现版本 [beringei](https://github.com/facebookincubator/beringei).



比较有趣的是, 国内最擅长多条腿走路的数据库及中间件巨头阿里在去年先后发布了自己的时序数据库产品 [TimescaleDB](https://yq.aliyun.com/articles/73537) 和 [HiTSDB](https://yq.aliyun.com/articles/162566) . 前者基于 PostgreSQL, 而后者同样基于 Gorilla, 并声称数据点平均内存开销 2 Byte 以下, 且具备 10M/sec 的数据点写入能力. 算一算, 和 Gorilla 论文中的数字相差不多.



##### prometheus/tsdb

[prometheus/tsdb](https://github.com/prometheus/tsdb) 是 prometheus 2.0 版本起使用的底层存储, 它的数据块编码也使用了 facebook 的 gorilla, 并具备了完整的持久化方案. 因此是我首先拿来参看的项目.



源码的阅读基于 [commit d45b59](https://github.com/prometheus/tsdb/tree/d45b595a1daefad23c09a2d994bf956f8b5f15a9) , 将会分成多篇文章.

 文章的数量, 篇幅和频率, 视乎我残存的悟性 😂.



以下是它的项目结构

```
├── Documentation
│   └── format
│       ├── chunks.md
│       ├── index.md
│       └── tombstones.md
├── LICENSE
├── README.md
├── block.go
├── block_test.go
├── chunkenc
│   ├── bstream.go
│   ├── chunk.go
│   ├── chunk_test.go
│   └── xor.go
├── chunks
│   └── chunks.go
├── cmd
│   └── tsdb
│       ├── Makefile
│       ├── README.md
│       └── main.go
├── compact.go
├── compact_test.go
├── db.go
├── db_test.go
├── encoding_helpers.go
├── fileutil
│   ├── dir_unix.go
│   ├── dir_windows.go
│   ├── fileutil.go
│   ├── mmap.go
│   ├── mmap_unix.go
│   ├── mmap_windows.go
│   ├── preallocate.go
│   ├── preallocate_darwin.go
│   ├── preallocate_linux.go
│   ├── preallocate_other.go
│   ├── sync.go
│   ├── sync_darwin.go
│   └── sync_linux.go
├── head.go
├── head_test.go
├── index
│   ├── encoding_helpers.go
│   ├── index.go
│   ├── index_test.go
│   ├── postings.go
│   └── postings_test.go
├── labels
│   ├── labels.go
│   ├── labels_test.go
│   └── selector.go
├── querier.go
├── querier_test.go
├── test
│   ├── conv_test.go
│   ├── hash_test.go
│   └── labels_test.go
├── testdata
│   └── 20kseries.json
├── testutil
│   └── testutil.go
├── tombstones.go
├── tombstones_test.go
├── tsdbutil
│   ├── buffer.go
│   └── buffer_test.go
├── wal.go
└── wal_test.go
```



##### 参考文章

关于时序数据库及常见产品的解读, 可以参考云栖社区的一系列文章:

[时间序列数据的存储和计算 - 概述](https://yq.aliyun.com/articles/104243?spm=5176.100239.0.0.1afb4721uSw5Gf)

[时间序列数据的存储和计算 - 开源时序数据库解析（一）](https://yq.aliyun.com/articles/104246?spm=5176.100239.0.0.1afb4721uSw5Gf)

[时间序列数据的存储和计算 - 开源时序数据库解析（二）](https://yq.aliyun.com/articles/106382?spm=5176.100239.0.0.1afb4721uSw5Gf)

[时间序列数据的存储和计算 - 开源时序数据库解析（三）](https://yq.aliyun.com/articles/158312?spm=5176.100239.0.0.1afb4721uSw5Gf)



