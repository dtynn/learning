package lec05

import (
	"math/rand"
)

type Random struct {
}

func (r Random) Select(in []int, k int) int {
	idx, kp := r.partition(in)
	if kp == k {
		return in[idx]
	}

	if kp > k {
		return r.Select(in[:idx], k)
	}

	return r.Select(in[idx+1:], k-kp)
}

func (r Random) partition(in []int) (int, int) {
	if len(in) == 1 {
		return 0, 1
	}

	idx := rand.Intn(len(in))
	pivot := in[idx]
	in[0], in[idx] = in[idx], in[0]

	i, j := 0, 1
	for ; j < len(in); j++ {
		if in[j] < pivot {
			i++
			in[i], in[j] = in[j], in[i]
		}
	}

	in[0], in[i] = in[i], in[0]
	return i, i + 1
}
