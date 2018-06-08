package add_binary

import (
	"testing"
)

func TestAddBinary(t *testing.T) {
	t.Run("0100+0000", func(t *testing.T) {
		a := "0100"
		b := "0000"
		t.Log(addBinary(a, b))
	})

	t.Run("+0000", func(t *testing.T) {
		a := ""
		b := "0000"
		t.Log(addBinary(a, b))
	})

	t.Run("0000+", func(t *testing.T) {
		a := "0000"
		b := ""
		t.Log(addBinary(a, b))
	})

	t.Run("1111+111", func(t *testing.T) {
		a := "1111"
		b := "111"
		t.Log(addBinary(a, b))
	})

	t.Run("1+111", func(t *testing.T) {
		a := "1"
		b := "111"
		t.Log(addBinary(a, b))
	})
}
