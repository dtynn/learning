package ch10

import (
	"reflect"
	"testing"

	"github.com/dtynn/learning/CLRS_BOOK/tests"
)

func TestSingleLinkList(t *testing.T) {
	doTest := func(t *testing.T, total int, l *SingleLinkList) {
		in := tests.GetDistinctInts(total)
		for _, i := range in {
			l.Insert(i)
		}

		head := l.Search(in[0])

		out := make([]int, total)
		out[0] = head.Key()

		for i := 1; i < total; i++ {
			head = head.Prev()
			out[i] = head.Key()
			l.Delete(head.Next())
		}

		if !reflect.DeepEqual(in, out) {
			t.Errorf("for in %v, got out %v", in, out)
		}
	}

	loop := 5
	l := NewSingleLinkList()

	t.Run("singleLinkList.5", func(t *testing.T) {
		for i := 0; i < loop; i++ {
			doTest(t, 5, l)
		}
	})

	t.Run("singleLinkList.25", func(t *testing.T) {
		for i := 0; i < loop; i++ {
			doTest(t, 25, l)
		}
	})

	t.Run("singleLinkList.50", func(t *testing.T) {
		for i := 0; i < loop; i++ {
			doTest(t, 50, l)
		}
	})
}
