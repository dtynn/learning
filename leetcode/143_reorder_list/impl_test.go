package impl

import (
	"bytes"
	"strconv"
	"testing"
)

func TestReorderList(t *testing.T) {
	cases := []struct {
		vals []int
	}{
		{
			vals: []int{1},
		},
		{
			vals: []int{1, 2, 3, 4, 5},
		},
		{
			vals: []int{1, 2, 3, 4, 5, 6},
		},
	}

	for i := range cases {
		head := makeList(cases[i].vals)
		printLinkList(t, head)
		reorderList(head)
		printLinkList(t, head)
	}
}

func makeList(vals []int) *ListNode {
	if len(vals) == 0 {
		return nil
	}

	first := &ListNode{
		Val: vals[0],
	}

	cur := first
	for i := range vals[1:] {
		cur.Next = &ListNode{
			Val: vals[1+i],
		}

		cur = cur.Next
	}

	return first
}

func printLinkList(t *testing.T, head *ListNode) {
	var buf bytes.Buffer
	buf.WriteString(strconv.Itoa(head.Val))

	cur := head.Next
	for cur != nil {
		buf.WriteString(", " + strconv.Itoa(cur.Val))
		cur = cur.Next
	}

	t.Log(buf.String())
}
