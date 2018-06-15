package maximal_reactangle

import (
	"testing"
)

func TestMaximalReactangle(t *testing.T) {
	t.Run("", func(t *testing.T) {
		matrix := [][]byte{
			{'1', '0', '1', '0', '0'},
			{'1', '0', '1', '1', '1'},
			{'1', '1', '1', '1', '1'},
			{'1', '0', '0', '1', '0'},
		}

		t.Log(maximalRectangle(matrix))
	})

	t.Run("2", func(t *testing.T) {
		matrix := [][]byte{
			{'1'},
		}

		t.Log(maximalRectangle(matrix))
	})

	t.Run("3", func(t *testing.T) {
		matrix := [][]byte{
			{'1', '1', '1', '1', '1', '1', '1', '1'},
			{'1', '1', '1', '1', '1', '1', '1', '0'},
			{'1', '1', '1', '1', '1', '1', '1', '0'},
			{'1', '1', '1', '1', '1', '0', '0', '0'},
			{'0', '1', '1', '1', '1', '0', '0', '0'},
		}

		t.Log(maximalRectangle(matrix))
	})

	t.Run("4", func(t *testing.T) {
		matrix := [][]byte{
			{'1', '0', '1', '1', '0', '1'},
			{'1', '1', '1', '1', '1', '1'},
			{'0', '1', '1', '0', '1', '1'},
			{'1', '1', '1', '0', '1', '0'},
			{'0', '1', '1', '1', '1', '1'},
			{'1', '1', '0', '1', '1', '1'},
		}

		t.Log(maximalRectangle(matrix))
	})
}
