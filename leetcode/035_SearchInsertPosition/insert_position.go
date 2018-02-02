package SearchInsertPosition

func searchInsert(nums []int, target int) int {
	i := 0
	for i < len(nums) {
		if nums[i] >= target {
			break
		}

		i++
	}

	return i
}
