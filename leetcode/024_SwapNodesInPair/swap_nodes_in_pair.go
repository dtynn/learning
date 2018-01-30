package SwapNodesInPair

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	var pprev *ListNode
	prev := head
	cur := head.Next

	loop := 0
	for cur != nil {
		if loop%2 == 0 {
			prev.Next = cur.Next
			cur.Next = prev

			if loop == 0 {
				head = cur
			} else {
				pprev.Next = cur
			}

			cur = prev
		}

		pprev = prev
		prev = cur
		cur = cur.Next

		loop++
	}

	return head
}
