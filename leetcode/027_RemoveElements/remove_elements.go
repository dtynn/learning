package RemoveElements

func removeElement(nums []int, val int) int {
	tail := len(nums)
	i := 0
	for i < len(nums) && i < tail {
		if nums[i] != val {
			i++
			continue
		}

		tail--
		nums[i], nums[tail] = nums[tail], nums[i]
	}

	return tail
}
