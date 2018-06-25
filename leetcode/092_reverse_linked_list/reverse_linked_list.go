package reverse_linked_list

// * Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if head == nil {
		return nil
	}

	if m == n {
		return head
	}

	count := n - m + 1

	first := head
	cur := head
	var part1Tail, revHead *ListNode

	num := 0

	for cur != nil {
		num++
		if num == m {
			revHead = reverse(cur, count)
			break
		}

		part1Tail = cur
		cur = cur.Next
	}

	if m == 1 {
		return revHead
	}

	part1Tail.Next = revHead

	return first
}

func reverse(head *ListNode, count int) *ListNode {
	if head == nil {
		return nil
	}

	var last, left *ListNode
	h := head
	cur := head
	reversed := 0
	for cur != nil {
		next := cur.Next
		cur.Next = last
		last = cur
		cur = next
		left = next
		reversed++
		if reversed >= count {
			cur = nil
		}
	}

	h.Next = left

	return last
}
