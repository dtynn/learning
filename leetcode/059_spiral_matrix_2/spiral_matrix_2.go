package spiral_matrix_2

func generateMatrix(n int) [][]int {
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}

	generate(1, matrix)
	return matrix
}

func generate(num int, matrix [][]int) {
	n := len(matrix)

	if n == 1 {
		matrix[0][0] = num
		return
	}

	rown, coln := 0, 0

	for coln < n {
		matrix[rown][coln] = num
		num++
		coln++
	}

	rown++
	coln--
	for rown < n {
		matrix[rown][coln] = num
		num++
		rown++
	}

	coln--
	rown--
	for coln >= 0 {
		matrix[rown][coln] = num
		num++
		coln--
	}

	rown--
	coln++
	for rown > 0 {
		matrix[rown][coln] = num
		num++
		rown--
	}

	if n == 2 {
		return
	}

	inner := make([][]int, 0, n-2)
	for line := 1; line < n-1; line++ {
		inner = append(inner, matrix[line][1:n-1])
	}

	generate(num, inner)
}
