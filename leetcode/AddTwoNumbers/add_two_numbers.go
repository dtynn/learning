package AddTwoNumbers

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	res := &ListNode{}
	current := res

	for {
		current.Val = current.Val + l1.Val + l2.Val

		if current.Val >= 10 {
			current.Val -= 10
			current.Next = &ListNode{
				Val: 1,
			}
		}

		l1 = l1.Next
		l2 = l2.Next
		if l1 == nil || l2 == nil {
			break
		}

		if current.Next == nil {
			current.Next = &ListNode{}
		}

		current = current.Next

	}

	if l1 != nil {
		if current.Next == nil {
			current.Next = l1
		} else {
			current.Next = addTwoNumbers(l1, current.Next)
		}
	} else if l2 != nil {
		if current.Next == nil {
			current.Next = l2
		} else {
			current.Next = addTwoNumbers(l2, current.Next)
		}
	}

	return res
}
