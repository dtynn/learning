package rotate_list

import (
	"testing"
)

func TestRoateList(t *testing.T) {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 5,
					},
				},
			},
		},
	}

	rotate := rotateRight(head, 6)
	for rotate != nil {
		t.Log(rotate.Val)
		rotate = rotate.Next
	}
}
