package permutations_2

func permuteUnique(nums []int) [][]int {
	if len(nums) == 0 {
		return nil
	}

	if len(nums) == 1 {
		return [][]int{
			{
				nums[0],
			},
		}
	}

	res := make([][]int, 0)
	exists := map[int]struct{}{}
	for i := range nums {
		head := nums[i]

		if _, ok := exists[head]; ok {
			continue
		}

		exists[head] = struct{}{}

		rest := make([]int, len(nums)-1)
		copy(rest[:i], nums[:i])
		copy(rest[i:], nums[i+1:])

		pers := permuteUnique(rest)
		for j := range pers {
			line := make([]int, len(nums))
			line[0] = head
			copy(line[1:], pers[j])

			res = append(res, line)
		}
	}

	return res
}
