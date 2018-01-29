package ThreeSumClosest

import (
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	if len(nums) < 3 {
		return 0
	}

	abs := func(n int) int {
		if n < 0 {
			return -n
		}

		return n
	}

	var resDiff int

	sort.Ints(nums)

	for i := range nums {
		m, n := i+1, len(nums)-1
		if i == 0 {
			resDiff = nums[0] + nums[1] + nums[2] - target
		}

		for n > m {
			total := nums[i] + nums[m] + nums[n]
			if total == target {
				return target
			}

			diff := total - target
			if abs(diff) < abs(resDiff) {
				resDiff = diff
			}

			if diff > 0 {
				n--
			} else {
				m++
			}
		}
	}

	return target + resDiff
}
