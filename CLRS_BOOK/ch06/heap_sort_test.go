package ch06

import (
	"testing"

	"github.com/dtynn/learning/CLRS_BOOK/tests"
)

func TestMaxHeapSort(t *testing.T) {
	tests.Sort(t, MaxHeap{}, 5)
}
