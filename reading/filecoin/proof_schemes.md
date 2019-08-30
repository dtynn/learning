### Proof-of-Replication and Proof-of-Spacetime

#### 动机

Filecoin 协议中, 存储提供者需要向提供存储证明, 并由区块链网络校验.



现有的证明方式, 如 *Provable Data Possesion (PDP)* 和 *Proof-of-Retrievability (PoR)* 只保证在校验当时,  证明人拥有部分数据. 而 Filecoin 需要更强的保证来阻止三类骗取奖励的攻击: 

1. Sybil attack

   通过创建多个 Sybil identities 来伪造保存的副本数量

2. outsourcing attacks

   通过从其他存储提供者处拉取数据来伪造自己能提供的存储容量

3. generation attacks

   没有实际存储数据, 而是通过一个体积小得多的程序按需生成数据内容, 来提高赢取区块奖励的可能性





#### Proof-of-Replication (PoRep)

定义:

1.  一种存储证明
2. 由证明方 (prover, P) 向校验方 (verifie, V)提供, 证明数据 (data D) 被复制到独立的物理存储空间
3. 一套交互协议:
   1. 提交物理上独立的数据 D 的副本数量 n
   2. 通过一套协议向 V 证明每个副本的独立性



由一组多项式时间的算法构成:

- PoRep.Setup $(1^\lambda , D)\rightarrow R, S_p, S_v$

  生成副本, 并向证明方和验证方提供必要的信息. 某些实现中要求一个第三方参与.

- PoRep.Prove $(S_p, R, c) \rightarrow \pi^c$

  - c 是由 V 提出的随机的校验要求
  - $\pi^c$ 是

- PoRep.Verify $(S_v, c, \pi^c) \rightarrow \{0, 1\}$

  检查证明的正确性



#### Proof-of-Spacetime

PoS 可以用来证明在校验当时, 数据的存在性和正确性.

但是需要有一种机制来确保数据被持续存储.

引入 Proof-of-Spacetime 来检查某段时间内数据被存储, 从直接上来说, 需要证明人:

1. 用生成 PoS 的序列方式来确定时间
2. 递归地执行生成过程来缩短证明的长度



与 PoRep 的算法构成相比, PoSt 引入了时间相关的参数



#### PoRep 和 PoSt 的实践

在 **Setup** 过程中引入一个缓慢的顺序计算过程 **Seal**.



##### 加密区块构建

- 避免碰撞的哈希
  - CRH
  - MerkleCRH: 将字符串拆分成构建二叉树, 递归地应用 CRH 并输出根节点信息
- zk-SNARKs (zero-knowledge Succinct Non-interactive ARguments of Knowledge)
  - 可靠方在初始化阶段生成两个公开的密钥: 用于证明过程的 pk 和用于校验过程的 vk
  - pk 用于证明人生成证明结果 $\pi$
  - $\pi$ 不需要额外的交互
  - vk 用于校验 $\pi$
  - zk-SNARK 的证明结果可以被公开校验, 校验人不需要与证明人交互
  - 证明结果具备常量大小, 并可以在线性时间内被校验
- 多项式时间算法