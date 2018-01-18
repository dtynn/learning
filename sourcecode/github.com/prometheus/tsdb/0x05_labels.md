#### labels

```
labels
├── labels.go
├── labels_test.go
└── selector.go
```

[labels](https://godoc.org/github.com/prometheus/tsdb/labels) 是标签, 对应 influxdb 中的 tags, 即一组键值对.

在 promethues/tsdb 中, timestamp 和 value 之外的所有信息都放在 labels 中

这个 pkg 的核心就是 Label, Labels, 以及 Labels 的 Matcher

##### Label

```go
// Label is a key/value pair of strings.
type Label struct {
	Name, Value string
}
```

就是键值对



##### Labels

```go
// Labels is a sorted set of labels. Order has to be guaranteed upon
// instantiation.
type Labels []Label
```

在实际使用中,  Labels 都应该是应该排序的. 因此 Labels 首先实现了 `sort.Interface`.

同时, Labels 之间也是可以进行比较的

```
// Compare compares the two label sets.
// The result will be 0 if a==b, <0 if a < b, and >0 if a > b.
func Compare(a, b Labels) int {
	l := len(a)
	if len(b) < l {
		l = len(b)
	}

	// 逐个 label 比较 name, value 的字母序
	for i := 0; i < l; i++ {
		if d := strings.Compare(a[i].Name, b[i].Name); d != 0 {
			return d
		}
		if d := strings.Compare(a[i].Value, b[i].Value); d != 0 {
			return d
		}
	}
	
	// If all labels so far were in common, the set with fewer labels comes first.
	// 可比较的部分无法确定顺序, 则比较两者长度
	return len(a) - len(b)
}
```



##### Slice

`Slice` 是 `Labels` 的切片

因为 `Labels` 可比较, 因此 `Slice` 也实现了 `sort.Interface`



##### Matcher

```
// Matcher specifies a constraint for the value of a label.
type Matcher interface {
	// Name returns the label name the matcher should apply to.
	Name() string
	// Matches checks whether a value fulfills the constraints.
	Matches(v string) bool
}
```

`Matcher` 用来筛选 Labels

这里提供了 equal, prefix, regexp, not 四种基本的 `Matcher`

