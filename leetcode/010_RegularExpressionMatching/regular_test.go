package RegularExpressionMatching

import (
	"testing"
)

func TestRegular(t *testing.T) {
	cases := []struct {
		s, p     string
		expected bool
	}{
		{
			s:        "",
			p:        "",
			expected: true,
		},
		{
			s:        "aa",
			p:        "a",
			expected: false,
		},
		{
			s:        "aa",
			p:        "aa",
			expected: true,
		},
		{
			s:        "aaa",
			p:        "aa",
			expected: false,
		},
		{
			s:        "aa",
			p:        "a*",
			expected: true,
		},
		{
			s:        "aa",
			p:        ".*",
			expected: true,
		},
		{
			s:        "ab",
			p:        ".*",
			expected: true,
		},
		{
			s:        "aab",
			p:        "c*a*b",
			expected: true,
		},
		{
			s:        "aaabbbb",
			p:        "a*abb*",
			expected: true,
		},
		{
			s:        "aaa",
			p:        "a*a*a*a*a*a*a",
			expected: true,
		},
		{
			s:        "",
			p:        ".*",
			expected: true,
		},
	}

	for i, c := range cases {
		if got := isMatch(c.s, c.p); got != c.expected {
			t.Errorf("#%d expected %v, got %v", i, c.expected, got)
		}
	}
}
