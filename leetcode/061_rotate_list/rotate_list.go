package rotate_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}

	nodes := make([]*ListNode, 0, 100)
	idx := 0
	for head != nil {
		nodes = append(nodes, head)
		idx++
		head = head.Next
	}

	k = k % len(nodes)
	if k == 0 {
		return nodes[0]
	}

	h := nodes[len(nodes)-k]
	prev := nodes[len(nodes)-k-1]
	last := nodes[len(nodes)-1]
	prev.Next = nil
	last.Next = nodes[0]
	return h
}
