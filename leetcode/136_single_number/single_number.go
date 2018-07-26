package single_number

func singleNumber(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	for i := 1; i < len(nums); i++ {
		nums[0] = nums[0] ^ nums[i]
	}

	return nums[0]
}
