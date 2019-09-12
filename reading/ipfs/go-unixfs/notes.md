### 0. Intro

#### 基本概念

[go-unixfs](https://github.com/ipfs/go-unixfs) 试图解决的是这样一个问题:

如何在分散的内容寻址数据之上构造一个类似文件系统的组织.



#### 预备知识

- [merkledag](https://github.com/ipfs/specs/tree/master/merkledag)
- [ipld](https://github.com/ipld/specs/)



### 1. 代码阅读

#### 1.1 项目结构

`go-unixfs` 的项目结构很简单:

```
./
├── LICENSE
├── README.md
├── file
│   └── unixfile.go
├── go.mod
├── go.sum
├── hamt // 实现了一套 Hash Array Mapped Trie
│   ├── hamt.go
│   ├── hamt_stress_test.go
│   ├── hamt_test.go
│   ├── util.go
│   └── util_test.go
├── importer // 用于将一般文件转换为 unixfs 文件
│   ├── balanced
│   │   ├── balanced_test.go
│   │   └── builder.go
│   ├── helpers
│   │   ├── dagbuilder.go
│   │   └── helpers.go
│   ├── importer.go
│   ├── importer_test.go
│   └── trickle
│       ├── trickle_test.go
│       └── trickledag.go
├── io // 提供 unixfs 上的文件读取和目录操作方法
│   ├── dagreader.go
│   ├── dagreader_test.go
│   ├── directory.go
│   ├── directory_test.go
│   ├── doc.go
│   └── resolve.go
├── mod // 实现了一个 DagModifier 来进行文件的写操作
│   ├── dagmodifier.go
│   └── dagmodifier_test.go
├── pb // protobuf 定义和生成文件
│   ├── Makefile
│   ├── unixfs.pb.go
│   └── unixfs.proto
├── test
│   └── utils.go
├── unixfs.go
└── unixfs_test.go

```



#### 1.2 Top Level

即 `./unixfs.go`, 定义了一些 unixfs 格式的数据结构, 并提供了相关的帮助函数



##### 1.2.1 数据类型

`go-unixfs` 中定义了以下几种文件系统对象

- Raw: 原始数据块
- File: 文件
- Directory: 目录
- Metadata: 元数据
- Symlink: 软链
- HAMTShard: hamt sharding



##### 1.2.2 帮助方法

首先是几类 fs 对象类型与`protobuf` 编码的二进制数据的相互转换:

- `func FromBytes(data []byte) (*pb.Data, error)`
- `func FilePBData(data []byte, totalsize uint64) []byte `
- `func FolderPBData() []byte `
- `func WrapData(b []byte) []byte `
- `func SymlinkData(path string) ([]byte, error) `
- `func HAMTShardData(data []byte, fanout uint64, hashType uint64) ([]byte, error) `
- `func UnwrapData(data []byte) ([]byte, error) `



对 `Raw` 和 `File` 类型的大小计算:

`func DataSize(data []byte) (uint64, error) `



##### 1.2.3 FSNode

提出一个 `FSNode` 的概念来表示 `go-unixfs` 定义下的文件系统对象, 并定义了 `Metadata` 结构体来记录`FSNode` 的附加信息:

```go
type FSNode struct {

	// UnixFS format defined as a protocol buffers message.
	format pb.Data
}

type Metadata struct {
	MimeType string
	Size     uint64
}

```



##### 1.2.4 FSNode 与 dag 节点的相互转换

- `func EmptyDirNode() *dag.ProtoNode ` 创建一个空目录节点
- `func ReadUnixFSNodeData(node ipld.Node) (data []byte, err error) ` 从一个 `ipld.Node` 中提取 `FSNode` 的编码结果数据
- `func ExtractFSNode(node ipld.Node) (*FSNode, error) ` 将`ipld.Node` 转换成`FSNode`



#### 1.3 