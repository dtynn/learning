package lec01

type Merge struct {
}

func (m Merge) Sort(in []int) []int {
	if len(in) == 1 {
		return in
	}

	return m.merge(m.Sort(in[:len(in)/2]), m.Sort(in[len(in)/2:]))
}

func (m Merge) merge(sorted1, sorted2 []int) []int {
	merged := make([]int, 0, len(sorted1)+len(sorted2))
	for len(sorted1) > 0 && len(sorted2) > 0 {
		if sorted1[0] < sorted2[0] {
			merged = append(merged, sorted1[0])
			sorted1 = sorted1[1:]
		} else {
			merged = append(merged, sorted2[0])
			sorted2 = sorted2[1:]
		}
	}

	merged = append(merged, sorted1...)
	merged = append(merged, sorted2...)
	return merged
}
