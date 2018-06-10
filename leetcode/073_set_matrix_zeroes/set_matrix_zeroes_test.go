package set_matrix_zeroes

import (
	"testing"
)

func TestSetZeroes(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		matrix := [][]int{
			{0, 1, 2, 0},
			{3, 4, 5, 2},
			{1, 3, 1, 5},
		}

		setZeroes(matrix)
		t.Log(matrix)
	})
}
