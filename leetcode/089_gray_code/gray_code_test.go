package gray_code

import (
	"testing"
)

func TestGrayCode(t *testing.T) {
	t.Run("0", func(t *testing.T) {
		n := 0
		t.Log(grayCode(n))
	})

	t.Run("1", func(t *testing.T) {
		n := 1
		t.Log(grayCode(n))
	})

	t.Run("2", func(t *testing.T) {
		n := 2
		t.Log(grayCode(n))
	})

	t.Run("3", func(t *testing.T) {
		n := 3
		t.Log(grayCode(n))
	})

	t.Run("4", func(t *testing.T) {
		n := 4
		t.Log(grayCode(n))
	})
}
