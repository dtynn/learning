package spiral_matrix

import (
	"testing"
)

func TestSpiralMatrix(t *testing.T) {
	t.Log(spiralOrder([][]int{
		{1},
	}))

	t.Log(spiralOrder([][]int{
		{1, 2},
		{4, 3},
	}))

	t.Log(spiralOrder([][]int{
		{1, 2, 3, 4},
		{10, 11, 12, 5},
		{9, 8, 7, 6},
	}))
}
