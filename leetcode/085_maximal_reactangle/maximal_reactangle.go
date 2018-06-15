package maximal_reactangle

func maximalRectangle(matrix [][]byte) int {
	n := len(matrix)
	if n == 0 {
		return 0
	}

	m := len(matrix[0])
	if m == 0 {
		return 0
	}

	widths := make([][]int, n)
	heights := make([][]int, n)
	for i := 0; i < n; i++ {
		widths[i] = make([]int, m)
		heights[i] = make([]int, m)
	}

	for row := range matrix {
		for col := range matrix[row] {
			if matrix[row][col] == '0' {
				continue
			}

			if widths[row][col] == 0 {
				w := width(matrix, m, row, col)
				for i := 0; i < w; i++ {
					widths[row][col+i] = w - i
				}
			}

			if heights[row][col] == 0 {
				h := height(matrix, n, row, col)
				for i := 0; i < h; i++ {
					heights[row+i][col] = h - i
				}
			}
		}
	}

	largest := 0

	for row := range widths {
		for col := range widths[row] {
			w := widths[row][col]
			if w == 0 {
				continue
			}

			minHeight := heights[row][col]
			if minHeight > largest {
				largest = minHeight
			}

			for i := 1; i < w; i++ {
				rightH := heights[row][col+i]
				if rightH < minHeight {
					minHeight = rightH
				}

				if area := minHeight * (i + 1); area > largest {
					largest = area
				}
			}
		}
	}

	return largest
}

func width(matrix [][]byte, m, row, col int) int {
	w := 0
	for i := col; i < m; i++ {
		if matrix[row][i] == '0' {
			break
		}

		w++
	}

	return w
}

func height(matrix [][]byte, n, row, col int) int {
	h := 0
	for i := row; i < n; i++ {
		if matrix[i][col] == '0' {
			break
		}

		h++
	}

	return h
}
