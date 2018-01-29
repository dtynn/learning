package ThreeSum

import "sort"

func threeSum(nums []int) [][]int {
	res := [][]int{}

	if len(nums) < 3 {
		return res
	}

	sort.Ints(nums)

	for i := range nums {
		if nums[i] > 0 {
			break
		}

		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		m, n := i+1, len(nums)-1
		exists := map[int]struct{}{}
		for n > m {
			if total := nums[i] + nums[m] + nums[n]; total == 0 {
				if _, ok := exists[nums[m]]; !ok {
					res = append(res, []int{nums[i], nums[m], nums[n]})
					exists[nums[m]] = struct{}{}
				}
				m++
				n--
			} else if total > 0 {
				n--
			} else {
				m++
			}
		}
	}

	return res
}
