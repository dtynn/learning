#### head.go



##### Head

Head 向调用方提供, 用于某个时间段内的数据读写.

Head 会同时处理 WAL 内的和已经持久化的数据.

Head 可以认为是`current Block`

所有 Block 不可再写入, Head 在写入有效期过后会转化为 Block 进行持久化.



###### Appender

```
// Appender returns a new Appender on the database.
// 会根据具体情形决定返回的 Appender 实例
// Appender 实例共两类
// initAppender 会在接收到第一个数据点时初始化 Head 的起始时间
// headAppender 逻辑相对简单
func (h *Head) Appender() Appender {
	h.metrics.activeAppenders.Inc()

	// The head cache might not have a starting point yet. The init appender
	// picks up the first appended timestamp as the base.
	if h.MinTime() == math.MinInt64 {
		return &initAppender{head: h}
	}
	return h.appender()
}

func (h *Head) appender() *headAppender {
	return &headAppender{
		head:          h,
		mint:          h.MaxTime() - h.chunkRange/2,
		samples:       h.getAppendBuffer(),
		highTimestamp: math.MinInt64,
	}
}
```



