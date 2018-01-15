#### tombstones.go

##### Stone

Stone 是作为删除数据的标记

```go
// Stone holds the information on the posting and time-range
// that is deleted.
type Stone struct {
	ref       uint64
	intervals Intervals
}
```



##### Interval, Intervals

用来记录时间段

```
// Interval represents a single time-interval.
type Interval struct {
	Mint, Maxt int64
}

func (tr Interval) inBounds(t int64) bool {
	return t >= tr.Mint && t <= tr.Maxt
}

func (tr Interval) isSubrange(dranges Intervals) bool {
	for _, r := range dranges {
		if r.inBounds(tr.Mint) && r.inBounds(tr.Maxt) {
			return true
		}
	}

	return false
}
```



##### TombstoneReader

```go
// TombstoneReader gives access to tombstone intervals by series reference.
type TombstoneReader interface {
	// Get returns deletion intervals for the series with the given reference.
	Get(ref uint64) (Intervals, error)

	// Iter calls the given function for each encountered interval.
	Iter(func(uint64, Intervals) error) error

	// Close any underlying resources
	Close() error
}
```



提供了一个内存版的实现

```go
type memTombstones map[uint64]Intervals

var emptyTombstoneReader = memTombstones{}

// EmptyTombstoneReader returns a TombstoneReader that is always empty.
func EmptyTombstoneReader() TombstoneReader {
	return emptyTombstoneReader
}

func (t memTombstones) Get(ref uint64) (Intervals, error) {
	return t[ref], nil
}

func (t memTombstones) Iter(f func(uint64, Intervals) error) error {
	for ref, ivs := range t {
		if err := f(ref, ivs); err != nil {
			return err
		}
	}
	return nil
}

func (t memTombstones) add(ref uint64, itv Interval) {
	t[ref] = t[ref].add(itv)
}

func (memTombstones) Close() error {
	return nil
}
```



TombstoneReader 的内容可以被写入文件, 也可以通过文件读出.

```go
func writeTombstoneFile(dir string, tr TombstoneReader) error {
	path := filepath.Join(dir, tombstoneFilename)
	tmp := path + ".tmp"
	
	// ...

	return renameFile(tmp, path)
}
```



```go
func readTombstones(dir string) (memTombstones, error) {
	b, err := ioutil.ReadFile(filepath.Join(dir, tombstoneFilename))
	// ...

	stonesMap := memTombstones{}

	for d.len() > 0 {
		// ...
		stonesMap.add(k, Interval{mint, maxt})
	}

	return stonesMap, nil
}
```



