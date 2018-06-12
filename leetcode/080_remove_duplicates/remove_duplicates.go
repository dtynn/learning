package remove_duplicates

func removeDuplicates(nums []int) int {
	loop := 0
	start := 0
	for loop < len(nums) {
		loop++

		if start-2 < 0 || nums[start] != nums[start-2] {
			start++
			continue
		}

		copy(nums[start:], nums[start+1:])
	}

	return start
}
