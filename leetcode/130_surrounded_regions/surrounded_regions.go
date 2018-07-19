package surrounded_regions

func solve(board [][]byte) {
	n := len(board)
	if n == 0 {
		return
	}

	m := len(board[0])
	if m == 0 {
		return
	}

	visited := map[[2]int]struct{}{}

	{
		row := 0
		for col := 0; col < m; col++ {
			if board[row][col] == 'X' {
				continue
			}

			c := [2]int{row, col}
			if _, ok := visited[c]; ok {
				continue
			}

			visitFromBoarder(board, visited, m, n, row, col)
		}
	}

	{
		row := n - 1
		for col := 0; col < m; col++ {
			if board[row][col] == 'X' {
				continue
			}

			c := [2]int{row, col}
			if _, ok := visited[c]; ok {
				continue
			}

			visitFromBoarder(board, visited, m, n, row, col)
		}
	}

	{
		col := 0
		for row := 1; row < n-1; row++ {
			if board[row][col] == 'X' {
				continue
			}

			c := [2]int{row, col}
			if _, ok := visited[c]; ok {
				continue
			}

			visitFromBoarder(board, visited, m, n, row, col)
		}
	}

	{
		col := m - 1
		for row := 1; row < n-1; row++ {
			if board[row][col] == 'X' {
				continue
			}

			c := [2]int{row, col}
			if _, ok := visited[c]; ok {
				continue
			}

			visitFromBoarder(board, visited, m, n, row, col)
		}
	}

	for row := 0; row < n; row++ {
		for col := 0; col < m; col++ {
			if board[row][col] == 'X' {
				continue
			}

			if _, ok := visited[[2]int{row, col}]; ok {
				continue
			}

			board[row][col] = 'X'
		}
	}
}

func visitFromBoarder(board [][]byte, visited map[[2]int]struct{}, m, n, row, col int) {
	if board[row][col] == 'X' {
		return
	}

	c := [2]int{row, col}
	if _, ok := visited[c]; ok {
		return
	}

	visited[c] = struct{}{}

	if row > 1 {
		visitFromBoarder(board, visited, m, n, row-1, col)
	}

	if row < n-2 {
		visitFromBoarder(board, visited, m, n, row+1, col)
	}

	if col > 1 {
		visitFromBoarder(board, visited, m, n, row, col-1)
	}

	if col < m-2 {
		visitFromBoarder(board, visited, m, n, row, col+1)
	}
}
