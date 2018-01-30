package FourSum

import (
	"sort"
)

func fourSum(nums []int, target int) [][]int {
	res := [][]int{}
	exists := map[[4]int]struct{}{}

	if len(nums) < 4 {
		return res
	}

	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		if i > 0 && (nums[i] == nums[i-1]) {
			continue
		}

		for j := i + 1; j < len(nums); j++ {
			m, n := j+1, len(nums)-1

			for n > m {
				total := nums[i] + nums[j] + nums[m] + nums[n]
				if total == target {
					a := [4]int{nums[i], nums[j], nums[m], nums[n]}
					if _, ok := exists[a]; !ok {
						res = append(res, []int{nums[i], nums[j], nums[m], nums[n]})
						exists[a] = struct{}{}
					}

					n--
					m++
					continue
				}

				if total > target {
					n--
				} else {
					m++
				}
			}
		}
	}

	return res
}
