package lec04

import (
	"math/rand"
)

type Quick struct {
}

func (q Quick) Sort(in []int) []int {
	if len(in) <= 1 {
		return in
	}

	ridx := rand.Intn(len(in))
	pivot := in[ridx]

	in[0], in[ridx] = in[ridx], in[0]

	i, j := 0, 1
	for ; j < len(in); j++ {
		if in[j] < pivot {
			i++
			in[i], in[j] = in[j], in[i]
		}
	}

	in[0], in[i] = in[i], in[0]
	q.Sort(in[:i])
	q.Sort(in[i+1:])

	return in
}
