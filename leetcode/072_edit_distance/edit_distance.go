package edit_distance

func minDistance(word1 string, word2 string) int {
	bw1 := []byte(word1)
	bw2 := []byte(word2)

	m := len(bw1)
	n := len(bw2)

	dp := make([][]int, 0, m+1)
	for i := 0; i < m+1; i++ {
		dp = append(dp, make([]int, n+1))
	}

	var min = func(n ...int) int {
		res := n[0]
		for i := 1; i < len(n); i++ {
			if n[i] < res {
				res = n[i]
			}
		}

		return res
	}

	for i := 0; i <= m; i++ {
		for j := 0; j <= n; j++ {
			var step int

			if i == 0 {
				step = j
			} else if j == 0 {
				step = i
			} else {
				if bw1[i-1] == bw2[j-1] {
					step = dp[i-1][j-1]
				} else {
					step = min(dp[i-1][j-1]+1, dp[i][j-1]+1, dp[i-1][j]+1)
				}
			}
			dp[i][j] = step
		}
	}

	return dp[m][n]
}
