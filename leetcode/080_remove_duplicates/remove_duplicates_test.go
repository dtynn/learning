package remove_duplicates

import (
	"testing"
)

func TestRemvoeDuplicates(t *testing.T) {
	t.Run("[1,1,1,2,2,3]", func(t *testing.T) {
		nums := []int{1, 1, 1, 2, 2, 3}
		length := removeDuplicates(nums)
		t.Log(length)
		t.Log(nums[:length])
	})
}
