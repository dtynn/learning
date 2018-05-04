package spiral_matrix

func spiralOrder(matrix [][]int) []int {
	order := make([]int, 0)
	spiral(matrix, &order)
	return order
}

func spiral(matrix [][]int, order *[]int) {
	rowCount := len(matrix)
	if rowCount == 0 {
		return
	}

	if rowCount == 1 {
		*order = append(*order, matrix[0]...)
		return
	}

	colCount := len(matrix[0])
	rown, coln := 0, 0

	for coln < colCount {
		*order = append(*order, matrix[rown][coln])
		coln++
	}

	coln--
	rown++
	for coln >= 0 && rown < rowCount {

		*order = append(*order, matrix[rown][coln])
		rown++
	}

	rown--
	coln--
	for rown >= 0 && coln >= 0 {
		*order = append(*order, matrix[rown][coln])
		coln--
	}

	rown--
	coln++
	for colCount > 1 && rown > 0 {
		*order = append(*order, matrix[rown][coln])
		rown--
	}

	if rowCount > 2 && colCount > 2 {
		inner := make([][]int, 0, rowCount-2)
		for r := 1; r < rowCount-1; r++ {
			row := make([]int, colCount-2)
			copy(row, matrix[r][1:])
			inner = append(inner, row)
		}

		spiral(inner, order)
	}
}
