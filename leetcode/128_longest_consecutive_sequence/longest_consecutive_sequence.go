package longest_consecutive_sequence

func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	exists := map[int]struct{}{}
	for i := range nums {
		exists[nums[i]] = struct{}{}
	}

	longest := 1
	for i := range nums {
		cur := nums[i]

		if _, ok := exists[cur]; !ok {
			continue
		}

		size := 1
		delete(exists, cur)

		larger := cur + 1
		smaller := cur - 1
		for {
			if _, ok := exists[larger]; !ok {
				break
			}

			delete(exists, larger)
			size++
			larger++
		}

		for {
			if _, ok := exists[smaller]; !ok {
				break
			}

			delete(exists, smaller)
			size++
			smaller--
		}

		if size > longest {
			longest = size
		}
	}

	return longest
}
