package rotate_image

func rotate(matrix [][]int) {
	n := len(matrix)

	for i := 0; i < n-1; i++ {
		// 4 x 4
		// [0, 0] => [0, 3]
		// [0, 3] => [3, 3]
		// [3, 3] => [3, 0]
		// [3, 0] => [0, 0]
		// [0, 1] => [1, 3] +1, +2
		// [1, 3] => [3, 2] +2, -1
		// [3, 2] => [2, 0] -1, -2
		// [2, 0] => [0, 1] -2, +1
		s1 := n - 1 - i
		s2 := i
		first := [2]int{0, i}
		second := [2]int{first[0] + s2, first[1] + s1}
		third := [2]int{second[0] + s1, second[1] - s2}
		forth := [2]int{third[0] - s2, third[1] - s1}

		matrix[first[0]][first[1]],
			matrix[second[0]][second[1]],
			matrix[third[0]][third[1]],
			matrix[forth[0]][forth[1]] = matrix[forth[0]][forth[1]],
			matrix[first[0]][first[1]], matrix[second[0]][second[1]], matrix[third[0]][third[1]]
	}

	if n > 3 {
		newMatrix := make([][]int, n-2)
		for line := 1; line < n-1; line++ {
			newMatrix[line-1] = matrix[line][1 : n-1]
		}

		rotate(newMatrix)
	}
}
