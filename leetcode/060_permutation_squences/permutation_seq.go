package permutation_squences

func getPermutation(n int, k int) string {
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = i + 1
	}

	seq := make([]byte, 0, n)

	for len(nums) > 0 {
		total := 1
		for i := n - 1; i > 0; i-- {
			total *= i
		}

		idx := (k - 1) / total
		if idx < 0 {
			idx = 0
		}

		// fmt.Println(nums, k, total, idx)

		seq = append(seq, '0'+byte(nums[idx]))
		nums = append(nums[:idx], nums[idx+1:]...)

		if k > total {
			k = k - idx*total
		}

		n--
	}

	return string(seq)
}
