package minimum_window_substring

import (
	"testing"
)

func TestSubstring(tt *testing.T) {
	tt.Run("BADEFGHBBBCBBEABBC_ABBC", func(tt *testing.T) {
		s := "BADEFGHBBBCBBEABBC"
		t := "ABBC"
		tt.Log(minWindow(s, t))
	})

	tt.Run("cabwefgewcwaefgcf_cae", func(tt *testing.T) {
		s := "cabwefgewcwaefgcf"
		t := "cae"
		tt.Log(minWindow(s, t))
	})
}
