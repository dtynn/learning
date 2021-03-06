package MergeKSortedList

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {

	return mergeLists(lists)
}

func mergeLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	if len(lists) == 1 {
		return lists[0]
	}

	if len(lists) == 2 {
		return mergeTwoLists(lists[0], lists[1])
	}

	return mergeTwoLists(mergeLists(lists[:len(lists)/2]), mergeLists(lists[len(lists)/2:]))
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	var head *ListNode
	if l1.Val < l2.Val {
		head = l1
		l1 = l1.Next
	} else {
		head = l2
		l2 = l2.Next
	}

	last := head
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			last.Next = l1
			l1 = l1.Next
		} else {
			last.Next = l2
			l2 = l2.Next

		}

		last = last.Next
	}

	if l1 == nil {
		last.Next = l2
	} else {
		last.Next = l1
	}

	return head
}
