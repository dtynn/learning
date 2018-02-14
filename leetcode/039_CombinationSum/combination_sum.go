package CombinationSum

import (
	"sort"
)

func combinationSum(candidates []int, target int) [][]int {
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
		*solutions = append(*solutions, solution)
		return
	}

	nextTarget := target - candidates[0]
	nextPrev := make([]int, len(prev)+1)
	copy(nextPrev, prev)
	nextPrev[len(prev)] = candidates[0]
	try(candidates, nextPrev, nextTarget, solutions)

	try(candidates[1:], prev, target, solutions)
}
