package unique_paths_2

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	n := len(obstacleGrid)
	if n == 0 {
		return 0
	}

	m := len(obstacleGrid[0])
	if m == 0 {
		return 0
	}

	sumGrid := make([][]int, n)
	for i := 0; i < n; i++ {
		sumGrid[i] = make([]int, m)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if obstacleGrid[i][j] == 1 {
				sumGrid[i][j] = 0
				continue
			}

			sum := 0
			if i == 0 && j == 0 {
				sum = 1
			} else if i == 0 {
				sum = sumGrid[0][j-1]
			} else if j == 0 {
				sum = sumGrid[i-1][0]
			} else {
				sum = sumGrid[i-1][j] + sumGrid[i][j-1]
			}

			sumGrid[i][j] = sum
		}
	}

	return sumGrid[n-1][m-1]
}
