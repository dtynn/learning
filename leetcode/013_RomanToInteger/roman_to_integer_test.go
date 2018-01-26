package RomanToInteger

import (
	"testing"
)

func TestRomanToInteger(t *testing.T) {
	cases := []struct {
		str string
		num int
	}{
		{
			str: "DCXXI",
			num: 621,
		},
	}

	for i, c := range cases {
		if got := romanToInt(c.str); got != c.num {
			t.Errorf("#%d expected %d, got %d", i+1, c.num, got)
		}
	}
}
