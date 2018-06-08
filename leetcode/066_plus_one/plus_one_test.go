package plus_one

import (
	"testing"
)

func TestPlusOne(t *testing.T) {
	t.Run("123", func(t *testing.T) {
		digits := []int{1, 2, 3}
		t.Log(plusOne(digits))
	})

	t.Run("4321", func(t *testing.T) {
		digits := []int{4, 3, 2, 1}
		t.Log(plusOne(digits))
	})

	t.Run("9", func(t *testing.T) {
		digits := []int{9}
		t.Log(plusOne(digits))
	})

	t.Run("99", func(t *testing.T) {
		digits := []int{9, 9}
		t.Log(plusOne(digits))
	})

	t.Run("199", func(t *testing.T) {
		digits := []int{1, 9, 9}
		t.Log(plusOne(digits))
	})

	t.Run("empty", func(t *testing.T) {
		digits := []int{}
		t.Log(plusOne(digits))
	})

	t.Run("000", func(t *testing.T) {
		digits := []int{0, 0, 0}
		t.Log(plusOne(digits))
	})

	t.Run("009", func(t *testing.T) {
		digits := []int{0, 0, 9}
		t.Log(plusOne(digits))
	})
}
