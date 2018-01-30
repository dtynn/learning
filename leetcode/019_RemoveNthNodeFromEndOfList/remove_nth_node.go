package RemoveNthNodeFromEndOfList

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	nodes := []*ListNode{head}
	total := 1
	cur := head.Next
	for cur != nil {
		nodes = append(nodes, cur)
		total++
		cur = cur.Next
	}

	nth := total - n
	if nth == 0 {
		return nodes[nth].Next
	}

	nodes[nth-1].Next = nodes[nth].Next
	return head
}
