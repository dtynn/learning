package subsets

import (
	"testing"
)

func TestSubsets(t *testing.T) {
	t.Run("[5, 1, 2, 3]", func(t *testing.T) {
		nums := []int{5, 1, 2, 3}
		t.Log(subsets(nums))
	})
}
