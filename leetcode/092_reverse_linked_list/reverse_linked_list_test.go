package reverse_linked_list

import (
	"bytes"
	"strconv"
	"testing"
)

func TestReverserLinkedList2(t *testing.T) {
	t.Run("[1, 2, 3, 4, 5], 2, 4", func(t *testing.T) {
		nums := []int{1, 2, 3, 4, 5}
		m := 2
		n := 4
		head := reverseBetween(linkedListFromInts(nums), m, n)
		printLinkedList(t, head)
	})
}

func linkedListFromInts(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}

	head := &ListNode{
		Val: nums[0],
	}

	cur := head
	for i := 1; i < len(nums); i++ {
		next := &ListNode{
			Val: nums[i],
		}

		cur.Next = next

		cur = next
	}

	return head
}

func printLinkedList(t *testing.T, head *ListNode) {
	var buf bytes.Buffer
	cur := head
	for cur != nil {
		buf.WriteString(strconv.Itoa(cur.Val))
		buf.WriteString(" => ")
		cur = cur.Next
	}

	buf.WriteString("nil")
	t.Log(buf.String())
}
