package ch10

import (
	"reflect"
	"testing"

	"github.com/dtynn/learning/CLRS_BOOK/tests"
)

func TestQueue(t *testing.T) {
	doTest := func(t *testing.T, total int, q *Queue) {
		in := tests.GetDistinctInts(total)
		for _, i := range in {
			err := q.Enqueue(i)
			if err != nil {
				t.Fatal(err)
			}
		}

		out := make([]int, total)
		for i := 0; i < total; i++ {
			x, err := q.Dequeue()
			if err != nil {
				t.Fatal(err)
			}

			out[i] = x
		}

		if !reflect.DeepEqual(in, out) {
			t.Errorf("for in %v, got out %v", in, out)
		}
	}

	loop := 5
	t.Run("queue.5", func(t *testing.T) {
		q := NewQueue(5)
		for i := 0; i < loop; i++ {
			doTest(t, 5, q)
		}
	})

	t.Run("queue.10", func(t *testing.T) {
		q := NewQueue(10)
		for i := 0; i < loop; i++ {
			doTest(t, 10, q)
		}
	})

	t.Run("queue.50", func(t *testing.T) {
		q := NewQueue(50)
		for i := 0; i < loop; i++ {
			doTest(t, 50, q)
		}
	})
}
