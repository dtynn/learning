### 5 Filecoin Storage and Retrieval Markets

- 两类具有同样结构但设计不同的市场:
  - Storage Market
  - Retrieval Market
- 客户和矿工都可以定价
- 交易所由网络运行, 而 Filecoin 中的所有节点构成网络
- 网络确保矿工提供服务换取奖励



#### 5.1 可校验的市场

交易市场是为特定货物或服务交易提供便利的协议.

在 Filecoin 这个场景中, 交易需要可以校验.

区中心网络的参与者可以校验供需双方的交易数据



我们提出"可校验市场"这个说法:

- 没有单一实体管理者
- 交易透明
- 任何人都可以匿名参与



用去中心化的方式进行交易操作:

参与者(矿工和全节点)独立校验所有相关内容



简述可校验市场的结构:

两阶段:

1. 订单匹配
   - 参与者添加 buy / sell 订单
   - 匹配的双方联合签订 deal 订单并提交到交易所, 并传播到网络上
2. 订单完成
   - 要求卖家生成加密证明来确保交易正常
   - 成功后, 网络进行支付和订单清理



#### 5.2 存储市场

满足数据存储需求的可校验市场



##### 5.2.1 要求

- 链上账本, 其重要性体现在:

  1. 矿工订单公开, 即价格透明
  2. 所有订单必须先进入区块链, 然后进入账单

- 参与者为资源提供担保

  - 避免违约
  - 存储矿工需要提供押金才能参与到存储市场中
  - 网络可以对矿工的违约行为进行做出惩罚

- 自治与异常处理

  - 只有当存储矿工可以在事先约定的数据生命周期内, 不断提供存储证明时, 订单才会终结
  - 网络必须能够校验证明的存在性和正确性, 并作出相应的处理.




##### 5.2.2 数据结构

- Put Orders:
  - bid: 来自客户
  - ask: 来自矿工
  - deal: 来自达成交易的双方
- Put Orderbook:
  - 当前合法的, 活跃的订单集合
  - 完全公开, 对所有诚实用户的展现一致
  - 当区块中出现新订单交易时, order 被添加到 book 中
  - 当订单取消, 过期或完成时, order 从 book 中移除



订单合法性的定义:

- bid: <size, funds[, price, time, collect, coding]>

  - 有足够的剩余资金
  - 时间
  - 确保满足一个最小的存储时长

- ask: <space, price>

  - 抵押物的数量和时效
  - 有效的空间大小

- deal: <ask, bid, ts>

  - ask: 指向一个唯一存在, 未被其他 deal order 引用过的 ask order

  - bid: 类似 ask , 指向 bid 类型的订单

  - ts: 确保在一个有效的时间区间

  - 如果恶意客户不将 deal order添加到 Orderbook, 则矿工的存储空间就会被一直占用.

    ts 可以避免这种情形出现

    ts 失效的订单是非法的, 不会被添加到 orderbook, 因此矿工可以在 ts 失效后重用被占用的存储空间



##### 5.2.3 存储市场协议

两个阶段:

- 订单匹配
  - 通过向区块链提交事务来添加订单
  - 匹配完成后, 进行数据块传输
  - 存储完成后, 双方签署并提交 deal order
- 订单完成
  - seal sector
  - 生成包含 piece 的 sector 的存储证明, 并提交到区块链
  - 网络的其他参与者必须校验证明



#### 5.3 数据提取市场

##### 5.3.1 要求

- 链下的 orderbook
  - 客户必须能够直接找到提供所需数据 piece 的矿工
  - 不通过区块链, 避免出现速度瓶颈
  - 参与者只需要持有部分 Orderbook
- 不需要可信的第三方介入
  - 在存储市场中, 区块链网络扮演了可信第三方的角色
  - 提取市场中没有这样一个角色
  - 通过分步传递数据, 分步支付来解决
  - 任何一方终止, 都不会继续造成损失
- 支付通道
  - 客户关注收到数据, 矿工关注收取费用
  - 通过公共账本来支付会成为瓶颈
  - 需要有效的链下支付方式
  - 只在解决纷争时使用区块链



##### 5.3.2 数据结构

- Get Order
- Get Orderbook



##### 5.3.3 提取市场协议

- 匹配

  - 私下沟通
  - 搭建小笔支付的通道

- 完成

  - 矿工发送小块数据, 收取客户端的收据
  - 矿工将收据提交到区块链上以获取奖励

  