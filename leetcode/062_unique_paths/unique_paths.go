package unique_paths

func uniquePaths(m int, n int) int {
	right := m - 1
	down := n - 1

	if down > right {
		right, down = down, right
	}

	res := 1
	for i := right + 1; i <= right+down; i++ {
		res = res * i
	}

	for i := down; i > 1; i-- {
		res = res / i
	}

	return res
}

func factorial(start, end int) int {
	res := 1
	for i := start; i <= end; i++ {
		if i != 0 {
			res *= i
		}
	}

	return res
}
