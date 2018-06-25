package subsets_2

import (
	"testing"
)

func TestSubsets2(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		nums := []int{1, 2, 2}
		t.Log(subsetsWithDup(nums))
	})

	t.Run("1", func(t *testing.T) {
		nums := []int{1, 2, 2, 2, 2, 3, 3, 4, 5}
		t.Log(subsetsWithDup(nums))
	})
}
