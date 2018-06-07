package valid_number

import (
	"testing"
)

func TestValidNumber(t *testing.T) {
	t.Run("0", func(t *testing.T) {
		s := "0"
		t.Log(isNumber(s))
	})

	t.Run(" 0.1 ", func(t *testing.T) {
		s := " 0.1 "
		t.Log(isNumber(s))
	})

	t.Run("2e10", func(t *testing.T) {
		s := "2e10"
		t.Log(isNumber(s))
	})

	t.Run("-.1", func(t *testing.T) {
		s := "-.1"
		t.Log(isNumber(s))
	})

	t.Run("-1.", func(t *testing.T) {
		s := "-1."
		t.Log(isNumber(s))
	})

	t.Run("-1.e-1", func(t *testing.T) {
		s := "-1.e-1"
		t.Log(isNumber(s))
	})

	t.Run("-1.e-005", func(t *testing.T) {
		s := "-1.e-005"
		t.Log(isNumber(s))
	})

	t.Run("11", func(t *testing.T) {
		s := "11"
		t.Log(isNumber(s))
	})

	t.Run("-001e-001", func(t *testing.T) {
		s := "-001e-001"
		t.Log(isNumber(s))
	})
}
