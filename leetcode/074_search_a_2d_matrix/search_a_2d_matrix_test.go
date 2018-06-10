package search_a_2d_matrix

import (
	"testing"
)

func TestSearch(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		matrix := [][]int{
			{1, 3, 5, 7},
			{10, 11, 16, 20},
			{23, 30, 34, 50},
		}

		target := 13

		t.Log(searchMatrix(matrix, target))
	})
}
