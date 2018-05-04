package insert_intervals

import (
	"sort"
)

// * Definition for an interval.
type Interval struct {
	Start int
	End   int
}

type Sorted []Interval

func (s Sorted) Len() int {
	return len(s)
}

func (s Sorted) Less(i, j int) bool {
	return s[i].Start < s[j].Start
}

func (s Sorted) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func insert(intervals []Interval, newInterval Interval) []Interval {
	return merge(append(intervals, newInterval))
}

func merge(intervals []Interval) []Interval {
	sort.Sort(Sorted(intervals))

	count := len(intervals)
	i := 1
	for i < count {
		cur := intervals[i]
		prev := intervals[i-1]
		if cur.Start > prev.End {
			i++
			continue
		}

		if cur.End > prev.End {
			intervals[i-1].End = cur.End
		}

		copy(intervals[i:], intervals[i+1:])
		count--
	}

	return intervals[:count]
}
