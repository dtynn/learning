package impl

// ListNode data structure
type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	pivot := head
	cur := head.Next
	pivot.Next = nil

	var leftHead, rightHead *ListNode
	var leftTail, rightTail *ListNode

	for cur != nil {
		next := cur.Next
		cur.Next = nil

		if cur.Val < pivot.Val {
			if leftTail == nil {
				leftHead = cur
				leftTail = cur
			} else {
				leftTail.Next = cur
				leftTail = cur
			}
		} else {
			if rightTail == nil {
				rightHead = cur
				rightTail = cur
			} else {
				rightTail.Next = cur
				rightTail = cur
			}
		}

		cur = next
	}

	lh := sortList(leftHead)
	rh := sortList(rightHead)
	pivot.Next = rh
	if lh == nil {
		return pivot
	}

	ltail := lh
	for ltail.Next != nil {
		ltail = ltail.Next
	}

	ltail.Next = pivot
	return lh
}
