package lec05

import (
	"math/rand"
	"testing"
)

func getTestArray(total int) []int {
	pool := make([]int, total)
	for i := 0; i < total; i++ {
		pool[i] = i + 1
	}

	res := make([]int, 0, total)
	for len(pool) > 0 {
		pick := rand.Intn(len(pool))
		res = append(res, pool[pick])
		pool = append(pool[:pick], pool[pick+1:]...)
	}

	return res
}

func TestSelect(t *testing.T) {
	kin100 := rand.Intn(100) + 1

	cases := []struct {
		in       []int
		k        int
		expected int
	}{
		{
			in:       []int{6, 10, 13, 5, 8, 3, 2, 11},
			k:        7,
			expected: 11,
		},
		{
			in:       getTestArray(100),
			k:        kin100,
			expected: kin100,
		},
	}

	r := Random{}

	for i, c := range cases {
		got := r.Select(c.in, c.k)
		if got != c.expected {
			t.Errorf("#%d get %d in %d, expected %d, got %d", i+1, c.k, len(c.in), c.expected, got)
		}
	}
}
