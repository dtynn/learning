package merge_sorted_array

import (
	"testing"
)

func TestMergeSortedArray(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		nums1 := []int{1, 2, 3, 0, 0, 0, 0}
		m := 3
		nums2 := []int{0, 2, 5, 6}
		n := 4

		merge(nums1, m, nums2, n)
		t.Log(nums1)
	})

	t.Run("2", func(t *testing.T) {
		nums1 := []int{1, 2, 3, 5, 6, 0, 0}
		m := 5
		nums2 := []int{0, 2}
		n := 2

		merge(nums1, m, nums2, n)
		t.Log(nums1)
	})

	t.Run("3", func(t *testing.T) {
		nums1 := []int{1, 2, 3, 0, 0, 0}
		m := 3
		nums2 := []int{2, 5, 6}
		n := 3

		merge(nums1, m, nums2, n)
		t.Log(nums1)
	})
}
