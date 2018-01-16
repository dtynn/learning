package TwoSum

func twoSum(nums []int, target int) []int {
	requires := map[int]int{}
	for i, n := range nums {
		if idx, ok := requires[n]; ok {
			return []int{idx, i}
		}

		requires[target-n] = i
	}

	return nil
}
