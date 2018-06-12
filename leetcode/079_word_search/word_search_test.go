package word_search

import (
	"testing"
)

func TestWordSearch(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		board := [][]byte{
			{'A', 'B', 'C', 'E'},
			{'S', 'F', 'C', 'S'},
			{'A', 'D', 'E', 'E'},
		}

		word := "BCCF"

		t.Log(exist(board, word))
	})

	t.Run("2", func(t *testing.T) {
		board := [][]byte{
			{'A', 'B', 'C', 'E'},
			{'S', 'F', 'E', 'S'},
			{'A', 'D', 'E', 'E'},
		}

		word := "ABCESEEEFS"

		t.Log(exist(board, word))
	})
}
