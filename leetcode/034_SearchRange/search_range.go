package SearchRange

func searchRange(nums []int, target int) []int {
	r := []int{-1, -1}
	search(r, nums, target, 0, len(nums))
	return r
}

func search(r []int, nums []int, target, head, tail int) {
	if tail == head {
		return
	}

	mid := head + (tail-head)/2
	if nums[mid] == target {
		r[0], r[1] = mid, mid

		for i := mid - 1; i >= head; i-- {
			if nums[i] == target {
				r[0] = i
			} else {
				break
			}
		}

		for j := mid + 1; j < tail; j++ {
			if nums[j] == target {
				r[1] = j
			} else {
				break
			}
		}
		return
	}

	if tail-head == 1 {
		return
	}

	if nums[mid] < target {
		search(r, nums, target, mid+1, tail)
	} else {
		search(r, nums, target, head, mid)
	}
}
