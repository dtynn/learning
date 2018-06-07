package minimum_path_sum

import (
	"testing"
)

func TestMinimumPathSum(t *testing.T) {
	t.Run("normal", func(t *testing.T) {
		grid := [][]int{
			{1, 3, 1},
			{1, 5, 1},
			{4, 2, 1},
		}

		t.Log(minPathSum(grid))
	})
}
