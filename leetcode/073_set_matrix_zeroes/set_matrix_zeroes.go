package set_matrix_zeroes

func setZeroes(matrix [][]int) {
	n := len(matrix)
	if n == 0 {
		panic(n)
	}

	m := len(matrix[0])
	if m == 0 {
		panic(m)
	}

	rows := make(map[int]struct{})
	cols := make(map[int]struct{})

	for r := 0; r < n; r++ {
		for c := 0; c < m; c++ {
			if matrix[r][c] == 0 {
				rows[r] = struct{}{}
				cols[c] = struct{}{}
			}
		}
	}

	for r := range rows {
		for i := 0; i < m; i++ {
			matrix[r][i] = 0
		}
	}

	for c := range cols {
		for i := 0; i < n; i++ {
			matrix[i][c] = 0
		}
	}
}
