#### chunkenc

```
chunkenc
├── bstream.go
├── chunk.go
├── chunk_test.go
└── xor.go
```



[chunkenc](https://godoc.org/github.com/prometheus/tsdb/chunkenc) 提供了时序数据点的编码格式.

它定义了一个 `Chunk` 接口 及其附属的 `Appender` 和 `Iterator`接口.

此外给出了 `Chunk` 的一个实现 `XORChunk` .



##### Chunk, Appender, Iterator

定义了数据块`Chunk`, 其为一组数据点的集合.

可以通过 `Appender` 继续写入, 及通过 `Iterator` 遍历已有的数据点.

这里明确指出了一个数据点是一对 时间戳 (int64) 和值 (float64).



##### Pool

`Chunk` 对象池的定义, 并给出了一个基于内存的实现.



##### XORChunk

这是目前给出的唯一一个 `Chunk` 实现, 使用了 Gorilla 的算法思路.

```
func (a *xorAppender) Append(t int64, v float64) {
	var tDelta uint64
	num := binary.BigEndian.Uint16(a.b.bytes())

	if num == 0 {
		// ...

	} else if num == 1 {
		// ...

	} else {
		tDelta = uint64(t - a.t)
		dod := int64(tDelta - a.tDelta)

		// Gorilla has a max resolution of seconds, Prometheus milliseconds.
		// Thus we use higher value range steps with larger bit size.
		switch {
		case dod == 0:
			a.b.writeBit(zero)
		case bitRange(dod, 14):
			a.b.writeBits(0x02, 2) // '10'
			a.b.writeBits(uint64(dod), 14)
		case bitRange(dod, 17):
			a.b.writeBits(0x06, 3) // '110'
			a.b.writeBits(uint64(dod), 17)
		case bitRange(dod, 20):
			a.b.writeBits(0x0e, 4) // '1110'
			a.b.writeBits(uint64(dod), 20)
		default:
			a.b.writeBits(0x0f, 4) // '1111'
			a.b.writeBits(uint64(dod), 64)
		}

		a.writeVDelta(v)
	}

	a.t = t
	a.v = v
	binary.BigEndian.PutUint16(a.b.bytes(), num+1)
	a.tDelta = tDelta
}
```



上述代码的 switch 部分对 dod (delta of delta) 的大小范围进行判定, 以确定一个最多 4bit 的标识, 及标识后的数据长度.

这里为了支持毫秒级的时间精度 (原始算法中为秒级), 对每一级的范围和长度做了调整.

数据点的压缩比率会受到一些影响, 但能适应更多的使用场景.



##### bstream

`XORChunk` 写入和读取点数据都依赖于 `bstream`提供的 bit 流读写能力, 核心是

```
func (b *bstream) writeBit(bit bit) {
	// ...
}

func (b *bstream) readBit() (bit, error) {
	// ...
}
```

```
func (b *bstream) writeByte(byt byte) {
	// ...
}

func (b *bstream) readByte() (byte, error) {
	// ...
}
```

```
func (b *bstream) writeBits(u uint64, nbits int) {
	// ...
}

func (b *bstream) readBits(nbits int) (uint64, error) {
	// ...
}
```

三组方法.