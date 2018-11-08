package impl

type ListNode struct {
	Val  int
	Next *ListNode
}

func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}

	count := 0
	cur := head
	for cur != nil {
		count++
		cur = cur.Next
	}

	var lHead, rHead *ListNode
	lHead = head
	cur = head
	for i := 0; i < count/2; i++ {
		rHead = cur.Next
		if i == count/2-1 {
			cur.Next = nil
		}

		cur = cur.Next
	}

	cur = head
	lHead = head.Next
	rHead = reverse(rHead)

	for lHead != nil || rHead != nil {
		if rHead != nil {
			cur.Next = rHead
			rHead = rHead.Next
			cur = cur.Next
		}

		if lHead != nil {
			cur.Next = lHead
			lHead = lHead.Next
			cur = cur.Next
		}
	}
}

func reverse(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var first *ListNode
	remain := head
	for remain != nil {
		next := remain.Next
		if first == nil {
			first = remain
			first.Next = nil
		} else {
			remain.Next = first
			first = remain
		}

		remain = next
	}

	return first
}
