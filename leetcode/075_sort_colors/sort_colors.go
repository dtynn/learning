package sort_colors

func sortColors(nums []int) {
	count := [3]int{}
	for i := range nums {
		count[nums[i]] = count[nums[i]] + 1
	}

	idx := 0
	for color, cnt := range count {
		for i := 0; i < cnt; i++ {
			nums[idx] = color
			idx++
		}
	}

}
