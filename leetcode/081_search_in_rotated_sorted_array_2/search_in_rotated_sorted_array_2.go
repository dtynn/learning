package search_in_rotated_sorted_array_2

func search(nums []int, target int) bool {
	if len(nums) == 0 {
		return false
	}

	if nums[0] == target {
		return true
	}

	if nums[0] < target {
		for i := 1; i < len(nums) && nums[i] >= nums[0]; i++ {
			if nums[i] == target {
				return true
			}

			if nums[i] > target {
				return false
			}
		}

		return false
	}

	for i := len(nums) - 1; i >= 1 && nums[i] <= nums[0]; i-- {
		if nums[i] == target {
			return true
		}

		if nums[i] < target {
			return false
		}
	}

	return false
}
