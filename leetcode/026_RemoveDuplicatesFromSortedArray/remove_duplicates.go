package RemoveDuplicatesFromSortedArray

func removeDuplicates(nums []int) int {
	i := 1
	for i < len(nums) {
		if nums[i] == nums[i-1] {
			nums = append(nums[:i], nums[i+1:]...)
			continue
		}

		i++
	}

	return len(nums)
}
