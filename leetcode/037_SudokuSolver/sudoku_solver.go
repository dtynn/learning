package SudokuSolver

// 先把所有格子的可能值都记录下来
// 找可能选择最少的格子
// 如果可能选择只有一个, 则填入, 更新其对应行, 列, 块其余格子的可能值, 再沿他的行, 列, 块尝试寻找下一个可以填入的格子
// 如果可能选择有多个, 则分别尝试填入
func solveSudoku(board [][]byte) {
	solve(board)
}

// return if finished
func solve(board [][]byte) bool {
	row, col := unfilled(board)
	if row == -1 && col == -1 {
		return true
	}

	var num byte
	for num = '1'; num <= '9'; num++ {
		if isValid(board, row, col, num) {
			board[row][col] = num
			if solve(board) {
				return true
			}

			board[row][col] = '.'
		}
	}

	return false
}

func unfilled(board [][]byte) (int, int) {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == '.' {
				return i, j
			}
		}
	}

	return -1, -1
}

func isValid(board [][]byte, row, col int, num byte) bool {
	for i := range board[row] {
		if i == col || board[row][i] == '.' {
			continue
		}

		if board[row][i] == num {
			return false
		}
	}

	for r := 0; r < len(board); r++ {
		if r == row || board[r][col] == '.' {
			continue
		}

		if board[r][col] == num {
			return false
		}
	}

	center := func(i int) int {
		switch i {
		case 0, 1, 2:
			return 1

		case 3, 4, 5:
			return 4

		default:
			return 7
		}
	}

	rowc := center(row)
	colc := center(col)
	for r := rowc - 1; r <= rowc+1; r++ {
		for c := colc - 1; c <= colc+1; c++ {
			if r == row && c == col {
				continue
			}

			if board[r][c] == '.' {
				continue
			}

			if board[r][c] == num {
				return false
			}
		}
	}

	return true
}
