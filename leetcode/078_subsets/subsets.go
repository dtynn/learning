package subsets

func subsets(nums []int) [][]int {
	prefix := make([]int, 0)
	res := make([][]int, 0)
	add(prefix, nums, &res)
	return res
}

func add(prefix, nums []int, res *[][]int) {
	cur := make([]int, len(prefix))
	copy(cur, prefix)
	*res = append(*res, cur)

	if len(nums) == 0 {
		return
	}

	for i := range nums {
		next := make([]int, len(prefix)+1)
		copy(next, prefix)
		next[len(next)-1] = nums[i]
		add(next, nums[i+1:], res)
	}
}
