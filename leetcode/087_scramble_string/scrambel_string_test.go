package scramble_string

import (
	"testing"
)

func TestScambleString(t *testing.T) {
	t.Run("1 false", func(t *testing.T) {
		s1 := "abcde"
		s2 := "caebd"
		t.Log(isScramble(s1, s2))
	})

	t.Run("2 true", func(t *testing.T) {
		s1 := "great"
		s2 := "tgrea"
		t.Log(isScramble(s1, s2))
	})

	t.Run("3 true", func(t *testing.T) {
		s1 := "a"
		s2 := "a"
		t.Log(isScramble(s1, s2))
	})

	t.Run("4 true", func(t *testing.T) {
		s1 := "bbaa"
		s2 := "abab"
		t.Log(isScramble(s1, s2))
	})

	t.Run("5 false", func(t *testing.T) {
		s1 := "abcd"
		s2 := "bdac"
		t.Log(isScramble(s1, s2))
	})

	t.Run("6 true", func(t *testing.T) {
		s1 := "abcd"
		s2 := "bcda"
		t.Log(isScramble(s1, s2))
	})

	t.Run("7 false", func(t *testing.T) {
		s1 := "abcdd"
		s2 := "dbdac"
		t.Log(isScramble(s1, s2))
	})

	t.Run("8", func(t *testing.T) {
		s1 := "ababcbaccbabbcbca"
		s2 := "bbbbbaaaacccccbba"
		t.Log(isScramble(s1, s2))
	})
}
