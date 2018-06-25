package decode_ways

import (
	"testing"
)

func TestDecodeWays(t *testing.T) {
	t.Run("12", func(t *testing.T) {
		s := "12"
		t.Log(numDecodings(s))
	})

	t.Run("102", func(t *testing.T) {
		s := "102"
		t.Log(numDecodings(s))
	})

	t.Run("279", func(t *testing.T) {
		s := "279"
		t.Log(numDecodings(s))
	})
}
