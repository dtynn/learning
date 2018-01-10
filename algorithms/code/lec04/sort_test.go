package lec04

import (
	"reflect"
	"testing"
)

func getInCopy(in []int) []int {
	res := make([]int, len(in))
	copy(res, in)
	return res
}

func TestSort(t *testing.T) {
	sorters := []interface {
		Sort(in []int) []int
	}{
		Quick{},
	}

	cases := []struct {
		in  []int
		out []int
	}{
		{
			in:  []int{8, 2, 4, 9, 3, 6},
			out: []int{2, 3, 4, 6, 8, 9},
		},
		{
			in:  []int{9, 8, 7, 6, 5, 4, 3, 2, 1},
			out: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			in:  []int{6, 10, 13, 5, 8, 3, 2, 11},
			out: []int{2, 3, 5, 6, 8, 10, 11, 13},
		},
	}

	for i := range sorters {
		t.Logf("sorter #%d %T", i+1, sorters[i])
		for j := range cases {
			out := sorters[i].Sort(getInCopy(cases[j].in))
			if !reflect.DeepEqual(out, cases[j].out) {
				t.Errorf("case #%d, expected %v, got %v", j+1, cases[j].out, out)
			}
		}
	}
}
