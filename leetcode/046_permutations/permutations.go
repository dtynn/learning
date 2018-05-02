package permutations

func permute(nums []int) [][]int {
	if len(nums) == 0 {
		return nil
	}

	if len(nums) == 1 {
		return [][]int{
			{nums[0]},
		}
	}

	res := make([][]int, 0)
	for i := range nums {
		head := nums[i]
		rest := make([]int, len(nums)-1)
		copy(rest[:i], nums[:i])
		copy(rest[i:], nums[i+1:])
		pers := permute(rest)
		for j := range pers {
			line := make([]int, len(nums))
			line[0] = head
			copy(line[1:], pers[j])
			res = append(res, line)
		}
	}

	return res
}
