package SearchInRotatedSortedArray

func search(nums []int, target int) int {
	size := len(nums)

	if size == 0 {
		return -1
	}

	i := 0
	for i > -len(nums) && i < len(nums) {
		idx := (i + size) % size
		if nums[idx] == target {
			return idx
		}

		direct := 1
		if nums[idx] > target {
			direct = -1
		}

		// 如果前进的方向与原方向相反, 则说明找不到
		if mul := direct * i; mul < 0 {
			break
		}

		i += direct
		// 如果下一个数增/减与预期不符, 则说明找不到
		next := (i + size) % size
		if (nums[next]-nums[idx])*direct < 0 {
			break
		}
	}

	return -1
}
