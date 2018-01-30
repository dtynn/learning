package ReverseNodesInKGroup

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil || k <= 1 {
		return head
	}

	space := make([]*ListNode, k)

	var lastTail *ListNode
	cur := head

	loop := 0
	for cur != nil {
		space[loop%k] = cur
		next := cur.Next

		// 进行 Reverse 操作
		if (loop+1)%k == 0 {
			// 第一轮, 需要修改 head
			if loop < k {
				head = cur
			}

			for i := 1; i < k; i++ {
				space[i].Next = space[i-1]
			}

			space[0].Next = next
			if lastTail != nil {
				lastTail.Next = cur
			}

			lastTail = space[0]
		}

		cur = next
		loop++
	}

	return head
}
