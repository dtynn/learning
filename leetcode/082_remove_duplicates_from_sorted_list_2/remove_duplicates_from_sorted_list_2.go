package remove_duplicates_from_sorted_list_2

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	h := &ListNode{
		Val:  head.Val - 1,
		Next: head,
	}

	last := h
	cur := head

	for cur != nil {
		if cur.Next == nil || cur.Next.Val != cur.Val {
			last = cur
			cur = cur.Next
			continue
		}

		next := cur.Next
		for next != nil {
			if next.Val != cur.Val {
				break
			}

			next = next.Next
		}

		last.Next = next
		cur = next
	}

	return h.Next
}
