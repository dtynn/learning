package NextPermutation

import "sort"

// 思路
// 长度 l, 遍历 l-1->0, 在[i+1:l] 中找到大于 nums[i] 的最小的数, 并将它换位到 i 上, 然后对 [i+1:l] 原址排序
func nextPermutation(nums []int) {
	i := len(nums) - 1 - 1
	for i >= 0 {
		target := nums[i]

		j := i + 1
		foundIdx := i
		for j < len(nums) {
			if nums[j] > target {
				if foundIdx == i || nums[j] < nums[foundIdx] {
					foundIdx = j
				}
			}
			j++
		}

		if foundIdx != i {
			nums[i], nums[foundIdx] = nums[foundIdx], nums[i]
			sort.Ints(nums[i+1:])
			return
		}
		i--
	}

	sort.Ints(nums)
}
