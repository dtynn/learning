package word_search

func exist(board [][]byte, word string) bool {
	if len(board) == 0 || len(board[0]) == 0 {
		return false
	}

	if len(word) == 0 {
		return false
	}

	n := len(board)
	m := len(board[0])

	bw := []byte(word)

	for row := range board {
		for col := range board[row] {
			if board[row][col] != bw[0] {
				continue
			}

			used := map[[2]int]struct{}{}

			start := [2]int{row, col}
			used[start] = struct{}{}

			if check(used, board, start, bw[1:], m, n) {
				return true
			}
		}
	}

	return false
}

func check(used map[[2]int]struct{}, board [][]byte, start [2]int, left []byte, m, n int) bool {
	if len(left) == 0 {
		return true
	}

	ns := neighbours(start, m, n)
	for _, next := range ns {
		b := board[next[0]][next[1]]
		if b != left[0] {
			continue
		}

		if _, ok := used[next]; ok {
			continue
		}

		used[next] = struct{}{}
		if check(used, board, next, left[1:], m, n) {
			return true
		}

		delete(used, next)
	}

	return false
}

func neighbours(pos [2]int, m, n int) [][2]int {
	ns := make([][2]int, 0, 4)
	maybe := [][2]int{
		{pos[0] - 1, pos[1]},
		{pos[0] + 1, pos[1]},
		{pos[0], pos[1] - 1},
		{pos[0], pos[1] + 1},
	}

	for i := 0; i < 4; i++ {
		p := maybe[i]
		if p[0] >= 0 && p[1] >= 0 && p[0] < n && p[1] < m {
			ns = append(ns, p)
		}
	}

	return ns
}
