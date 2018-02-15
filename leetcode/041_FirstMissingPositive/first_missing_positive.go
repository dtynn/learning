package FirstMissingPositive

// 长度为 l 的数组, 第一个缺失的正数必然在 [1, l+1] 之间
func firstMissingPositive(nums []int) int {
	exists := make([]bool, len(nums))
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 && nums[i] <= len(nums) {
			exists[nums[i]-1] = true
		}
	}

	for n := 0; n < len(exists); n++ {
		if !exists[n] {
			return n + 1
		}
	}

	return len(exists) + 1
}
