package remove_duplicates_from_sorted_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	h := head

	last := h
	cur := head.Next

	for cur != nil {
		if cur.Val == last.Val {
			last.Next = cur.Next
		} else {
			last = cur
		}

		cur = cur.Next
	}

	return h
}
