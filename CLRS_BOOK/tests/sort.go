package tests

import (
	"math/rand"
	"sort"
	"testing"
)

func ReverseInts(src []int) {
	for i := 0; i < len(src)/2; i++ {
		src[i], src[len(src)-1-i] = src[len(src)-1-i], src[i]
	}
}

func CopyInts(src []int) []int {
	dst := make([]int, len(src))
	copy(dst, src)
	return dst
}

func GetDistinctInts(total int) []int {
	ints := make([]int, total)
	for i := 0; i < total; i++ {
		ints[i] = i
	}

	res := ints[:]
	var ir int

	for len(ints) > 0 {
		idx := rand.Intn(len(ints))
		res[ir], ints[idx] = ints[idx], res[ir]
		ir++
		ints = ints[1:]
	}

	return res
}

type Sorter interface {
	Sort([]int) []int
}

func Sort(t *testing.T, sorter Sorter, loop int) {
	doTest := func(t *testing.T, total int) {
		for i := 0; i < loop; i++ {
			ints := GetDistinctInts(total)
			sorted := sorter.Sort(CopyInts(ints))
			if !sort.IntsAreSorted(sorted) {
				t.Errorf("for case %v, get unsorted result %v", ints, sorted)
			}
		}
	}

	t.Run("sort.0", func(t *testing.T) {
		doTest(t, 0)
	})

	t.Run("sort.5", func(t *testing.T) {
		doTest(t, 5)
	})

	t.Run("sort.10", func(t *testing.T) {
		doTest(t, 10)
	})

	t.Run("sort.100", func(t *testing.T) {
		doTest(t, 100)
	})
}
