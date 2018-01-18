package ch10

import (
	"reflect"
	"testing"

	"github.com/dtynn/learning/CLRS_BOOK/tests"
)

func TestStack(t *testing.T) {
	doTest := func(t *testing.T, total int, s *Stack) {
		in := tests.GetDistinctInts(total)
		for _, i := range in {
			if err := s.Push(i); err != nil {
				t.Fatal(err)
			}
		}

		out := make([]int, total)
		for i := 0; i < total; i++ {
			x, err := s.Pop()
			if err != nil {
				t.Fatal(err)
			}

			out[i] = x
		}

		rev := tests.CopyInts(out)
		tests.ReverseInts(rev)
		if !reflect.DeepEqual(in, rev) {
			t.Errorf("for in %v, got out %v", in, out)
		}
	}

	loop := 5
	t.Run("stack.5", func(t *testing.T) {
		s := NewStack(5)
		for i := 0; i < loop; i++ {
			doTest(t, 5, s)
		}
	})

	t.Run("stack.10", func(t *testing.T) {
		s := NewStack(10)
		for i := 0; i < loop; i++ {
			doTest(t, 10, s)
		}
	})

	t.Run("stack.50", func(t *testing.T) {
		s := NewStack(50)
		for i := 0; i < loop; i++ {
			doTest(t, 50, s)
		}
	})
}
