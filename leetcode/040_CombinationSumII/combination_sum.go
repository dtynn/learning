package CombinationSumII

import "sort"

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	solutions := make([][]int, 0)
	try(candidates, nil, target, &solutions)
	return solutions
}

func try(candidates, prev []int, target int, solutions *[][]int) {
	// 如果没有剩余元素, 或者剩余元素都比目标大
	if len(candidates) == 0 || candidates[0] > target {
		return
	}

	// 剩余的第一个元素恰好等于目标
	if candidates[0] == target {
		solution := make([]int, len(prev)+1)
		copy(solution, prev)
		solution[len(prev)] = candidates[0]

		exists := *solutions
		for i := 0; i < len(exists); i++ {
			if eqaul(solution, exists[i]) {
				return
			}
		}

		*solutions = append(exists, solution)
		return
	}

	nextTarget := target - candidates[0]
	nextPrev := make([]int, len(prev)+1)
	copy(nextPrev, prev)
	nextPrev[len(prev)] = candidates[0]
	try(candidates[1:], nextPrev, nextTarget, solutions)

	try(candidates[1:], prev, target, solutions)
}

func eqaul(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
