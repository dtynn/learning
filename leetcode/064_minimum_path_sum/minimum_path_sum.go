package minimum_path_sum

func minPathSum(grid [][]int) int {
	n := len(grid)
	if n == 0 {
		return 0
	}

	m := len(grid[0])
	if m == 0 {
		return 0
	}

	minGrid := make([][]int, n)
	for i := 0; i < n; i++ {
		minGrid[i] = make([]int, m)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			min := 0
			if i == 0 && j == 0 {
				min = 0
			} else if i == 0 {
				min = minGrid[0][j-1]
			} else if j == 0 {
				min = minGrid[i-1][0]
			} else {
				min = minGrid[i-1][j]
				if minGrid[i][j-1] < min {
					min = minGrid[i][j-1]
				}
			}

			minGrid[i][j] = min + grid[i][j]
		}
	}

	return minGrid[n-1][m-1]
}
