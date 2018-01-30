package ValidParentheses

import "testing"

func TestValidParentheses(t *testing.T) {
	cases := []struct {
		s     string
		valid bool
	}{
		{
			s:     "",
			valid: true,
		},
		{
			s:     "{",
			valid: false,
		},
		{
			s:     "}",
			valid: false,
		},
		{
			s:     "[{]",
			valid: false,
		},
		{
			s:     "[}]",
			valid: false,
		},
		{
			s:     "[]{",
			valid: false,
		},
		{
			s:     "[]}",
			valid: false,
		},
		{
			s:     "[]",
			valid: true,
		},
		{
			s:     "[]()",
			valid: true,
		},
		{
			s:     "[()]",
			valid: true,
		},
		{
			s:     "[[[]",
			valid: false,
		},
	}

	for _, c := range cases {
		if c.valid != isValid(c.s) {
			t.Errorf("for %s, got %v", c.s, !c.valid)
		}
	}
}
