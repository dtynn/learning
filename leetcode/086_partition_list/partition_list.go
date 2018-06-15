package partition_list

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	if head == nil {
		return nil
	}

	ltHead := &ListNode{}
	ltTail := ltHead

	gteHead := &ListNode{}
	gteTail := gteHead

	cur := head
	for cur != nil {
		next := cur.Next

		if cur.Val < x {
			ltTail.Next = cur
			ltTail = cur
			ltTail.Next = nil
		} else {
			gteTail.Next = cur
			gteTail = cur
			gteTail.Next = nil
		}

		cur = next
	}

	ltTail.Next = gteHead.Next
	return ltHead.Next
}
