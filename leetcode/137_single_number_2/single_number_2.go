package single_number_2

func singleNumber(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	var ones, twos int
	for i := 0; i < len(nums); i++ {
		ones = (ones ^ nums[i]) & (^twos)
		twos = (twos ^ nums[i]) & (^ones)
	}

	return ones
}
