package unique_binary_search_tree

func numTrees(n int) int {
	if n == 0 {
		return 0
	}

	return countTrees(n)
}

func countTrees(total int) int {
	if total == 0 {
		return 1
	}

	if total == 1 {
		return 1
	}

	if total == 2 {
		return 2
	}

	res := 0

	for leftNum := 0; leftNum < total; leftNum++ {
		rightNum := total - 1 - leftNum
		res += countTrees(leftNum) * countTrees(rightNum)
	}

	return res
}
