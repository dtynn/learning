package search_a_2d_matrix

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	tryRow := -1

	for row := 0; row < len(matrix); row++ {
		head := matrix[row][0]
		if head == target {
			return true
		}

		if head > target {
			break
		}

		tryRow = row
	}

	if tryRow < 0 {
		return false
	}

	return func(r []int) bool {
		for _, i := range r {
			if i == target {
				return true
			}
		}

		return false
	}(matrix[tryRow])
}
