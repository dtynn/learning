package maximum_subarray

func maxSubArray(nums []int) int {
	last := nums[0]
	max := last
	for i := 1; i < len(nums); i++ {
		cur := nums[i]
		if last > 0 {
			cur += last
		}

		if cur > max {
			max = cur
		}

		last = cur
	}

	return max
}
