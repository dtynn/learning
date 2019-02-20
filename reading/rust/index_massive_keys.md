## Index 1,600,000,000 Keys with Automata and Rust

[原文](https://blog.burntsushi.net/transducers/)

使用有限状态机来存储有序字符串集合或map: 存储高效, 检索快速.

文章概述:

- 使用有限状态机作为实现有序集合或 map 的数据结构
  - 一个 rust 实现的 lib: [fst](https://github.com/BurntSushi/fst)
  - 通过命令行使用的范例
  - 一些测试和试验
- 其他:
  - memory maps
  - automaton intersection with regular expressions
  - fuzzy searching with Levenshtein distance and streaming set operations



### Teaser

使用 ~16,000,000 Wikipedia article titles 的数据集,

- 文件体积与 gzip 压缩结果的对比
- 模糊搜索与 grep 查询结果的对比



### Table of Contents

1. fsm 及其应用的抽象描述

   介绍 fsm 的心智模型

2. 对 1 中的抽象提供一个范例实现

   - 如何使用 fst lib
   - 一些实现细节

3. 使用命令行工具创建索引的范例

   - 使用真实数据集
   - 性能分析



### Finite state machines as data structures

FSM: 状态和状态转移的集合.

一个 FSM 同一时间只有一个状态



#### Ordered sets

使用 *deterministic acyclic finite state acceptor* 实现有序集合

*deterministic acyclic finite state acceptor* 是满足以下特性的 FSM:

1. Deterministic 确定性: 在任意给定状态,  对任意输入, 最多只有一个可选路径.

2. Acyclic 非循环: 对于已经访问过的状态, 无法再次访问.

3. An acceptor 受体: 对于一组特定的输入, 当且仅当fsm 在输入尾端处在 "final" 状态时, 我们才说该 fsm "接受" 此输入.

   这个标准在有序 map 的场景下会有变化.



搜索某个 key 的耗时与 keys 的总数无关.

集合的遍历方式.



#### Ordered maps

使用 *deterministic acyclic finite state transducer* 实现有序 map

1. Deterministic

2. Acyclic

3. A **transducer**: 

   在 FSM 中:

   - 每个存在于其中的 输入 对应一个值.
   - 当且仅当某个输入使 FSM 处于 final 状态时, 才有值与其对应



#### Construction

有序 set 和 map 中的元素需要按照词典序(lexicographic order) 添加



##### Trie construction

- trie 与 fsa 性质相近
- trie 与 fsa 的差别在于 trie 仅允许共享前缀, 而 fsa 也允许共享后缀.



##### FSA construction

- 构造一个 trie, 再对其使用 [general minimization algorithm](https://en.wikipedia.org/wiki/DFA_minimization) , 可以达成共享后缀的目标.

  - general minimization algorithms 可能会有很高的时间和空间复杂度.

  - 当以词典序添加元素时, fsa 中, 与新增元素没有共享前缀的部分可以被冻结.

    这是因为后续所有的元素都一定比当前新增的元素排序靠后, 也就更不可能与冻结的部分有共享前缀.



##### FST construction

- 与 FSA 大致相同
- 最主要的差别是路径上的值的位置, 和共享



##### Construction in practice

通过提供一个额外的 hash 表, 我们可以将检测某个新元素中可重用部分的时间大幅缩减.

对使用排好序的元素构造一个近似最小的 FST, 我们可以将复杂度控制在线性时间, 常数空间.



##### References



### The FST Library



#### Building ordered sets and maps



#### Querying ordered sets and maps



#### Memory maps



#### Levenshtein automata



#### Levenshtein automata and Unicode



#### Regular expressions



#### Set operations



#### Raw transducers



### The FST command line tool



#### How to get it



#### Brief introduction



#### Experiments



### Lessons and trade offs

- Not a general purpose data structure
- More useful on 64 bit systems
- Requires fast random access



### Conclusion

