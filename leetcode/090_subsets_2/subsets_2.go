package subsets_2

import (
	"sort"
)

func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	res := make([][]int, 0)
	appendSubsets([]int{}, nums, &res)
	return res
}

func appendSubsets(prefix, left []int, res *[][]int) {
	*res = append(*res, prefix)
	if len(left) == 0 {
		return
	}

	for i := 0; i < len(left); i++ {
		if i > 0 && left[i] == left[i-1] {
			continue
		}

		nextPrefix := make([]int, len(prefix)+1)
		copy(nextPrefix, prefix)
		nextPrefix[len(nextPrefix)-1] = left[i]
		nextLeft := make([]int, len(left)-i-1)
		copy(nextLeft, left[i+1:])
		appendSubsets(nextPrefix, nextLeft, res)
	}
}
