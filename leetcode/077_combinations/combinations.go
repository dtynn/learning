package combinations

func combine(n int, k int) [][]int {
	if n < k {
		return nil
	}

	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = i + 1
	}

	if n == k {
		return [][]int{nums}
	}

	res := make([][]int, 0)
	add(k, nil, nums, &res)
	return res
}

func add(k int, prefix, nums []int, res *[][]int) {
	for i := range nums {
		nextPrefix := make([]int, len(prefix)+1)
		copy(nextPrefix, prefix)
		nextPrefix[len(nextPrefix)-1] = nums[i]
		if len(nextPrefix) == k {
			*res = append(*res, nextPrefix)
			continue
		}

		add(k, nextPrefix, nums[i+1:], res)
	}
}
