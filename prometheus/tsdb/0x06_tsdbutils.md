#### tsdbutil

```
tsdbutil
├── buffer.go
└── buffer_test.go
```

[tsdbutil](https://godoc.org/github.com/prometheus/tsdb/tsdbutil) 目前只提供了一个 `BufferedSeriesIterator`

##### sampleRing

```go
type sample struct {
	t int64
	v float64
}

// 既然是 Ring, 那么 buf 就是环装的, 因此有辅助的 i, f, l
type sampleRing struct {
	delta int64

	buf []sample // lookback buffer
	i   int      // position of most recent element in ring buffer
	f   int      // position of first element in ring buffer
	l   int      // number of elements in buffer
}
```

sampleRing 用来处理数据点的采样

```go
// add adds a sample to the ring buffer and frees all samples that fall
// out of the delta range.
func (r *sampleRing) add(t int64, v float64) {
	l := len(r.buf)
	// Grow the ring buffer if it fits no more elements.
	if l == r.l {
		// ring buffer 的扩容
		buf := make([]sample, 2*l)
		copy(buf[l+r.f:], r.buf[r.f:])
		copy(buf, r.buf[:r.f])

		r.buf = buf
		r.i = r.f
		r.f += l
	} else {
		r.i++
		if r.i >= l {
			r.i -= l
		}
	}

	r.buf[r.i] = sample{t: t, v: v}
	r.l++

	// Free head of the buffer of samples that just fell out of the range.
	// 这里认为 add 是有序的, 将头部所有早于 `t - r.delta` 的数据点移出有效区域
	for r.buf[r.f].t < t-r.delta {
		r.f++
		if r.f >= l {
			r.f -= l
		}
		r.l--
	}
}
```



##### sampleRingIterator

```
type sampleRingIterator struct {
	r *sampleRing
	i int
}
```

sampleRingIterator 是 [SeriesIterator](https://github.com/prometheus/tsdb/blob/d45b595a1daefad23c09a2d994bf956f8b5f15a9/querier.go#L668-L680) 的实现



##### BufferedSeriesIterator

BufferedSeriesIterator 同样也实现了 [SeriesIterator](https://github.com/prometheus/tsdb/blob/d45b595a1daefad23c09a2d994bf956f8b5f15a9/querier.go#L668-L680), 它将一段部分数据点通过 sampleRing 缓存下来, 具体效果, 待阅读其他代码.

```go
// BufferedSeriesIterator wraps an iterator with a look-back buffer.
type BufferedSeriesIterator struct {
	it  tsdb.SeriesIterator
	buf *sampleRing

	lastTime int64
}

// NewBuffer returns a new iterator that buffers the values within the time range
// of the current element and the duration of delta before.
// BufferedSeriesIterator 的作用是对上层 Iter 进行封装
// 将其中最多 delta 时间范围内的数据点通过 sampleRing 缓存下来
func NewBuffer(it tsdb.SeriesIterator, delta int64) *BufferedSeriesIterator {
	return &BufferedSeriesIterator{
		it:       it,
		buf:      newSampleRing(delta, 16),
		lastTime: math.MinInt64,
	}
}
```



###### Seek

```
// Seek advances the iterator to the element at time t or greater.
// 这里的 `指针` 只会向后移动, 不会向前
func (b *BufferedSeriesIterator) Seek(t int64) bool {
	t0 := t - b.buf.delta

	// If the delta would cause us to seek backwards, preserve the buffer
	// and just continue regular advancement while filling the buffer on the way.
	// 此时 sampleRing 中的点都会失效, 因此直接重置
	if t0 > b.lastTime {
		b.buf.reset()

		ok := b.it.Seek(t0)
		if !ok {
			return false
		}
		b.lastTime, _ = b.At()
	}

	if b.lastTime >= t {
		return true
	}
	for b.Next() {
		if b.lastTime >= t {
			return true
		}
	}

	return false
}
```

