package distinct_subsequences

func numDistinct(s string, t string) int {
	if s == t {
		return 1
	}

	m := len(s)
	n := len(t)

	if n == 0 {
		return 1
	}

	if m < n {
		return 0
	}

	nums := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		nums[i] = make([]int, m+1)
		if i == 0 {
			for j := 0; j < m+1; j++ {
				nums[i][j] = 1
			}
		}
	}

	for i := 1; i < n+1; i++ {
		for j := 1; j < m+1; j++ {
			num := nums[i][j-1]
			if s[j-1] == t[i-1] {
				num += nums[i-1][j-1]
			}

			nums[i][j] = num
		}
	}

	return nums[n][m]
}
