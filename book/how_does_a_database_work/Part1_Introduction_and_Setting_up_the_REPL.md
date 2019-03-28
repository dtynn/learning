### Part 1 - Introduction and Setting up the REPL



#### Sqlite

##### architecture

```
Tokenizer => Parser => Code Generator => Virtual Machine => B-Tree => Pager => OS Interface
```



##### front-end

- tokenizer
- paser
- code generator

输入: SQL Query

输出: VM 的 指令

##### back-end

- virtual machine: 执行前端传递过来的指令
- B-tree: 数据页的索引
- pager: 响应数据页的读写操作
- os interface: 与操作系统的交互层



#### Making a Simple REPL

