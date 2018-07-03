package convert_sorted_list_to_bst

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}

	nums := make([]int, 0)
	cur := head
	for cur != nil {
		nums = append(nums, cur.Val)
		cur = cur.Next
	}

	return sortedArrayToBST(nums)
}

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	size := len(nums)
	mid := size / 2

	root := &TreeNode{
		Val: nums[mid],
	}

	root.Left = sortedArrayToBST(nums[:mid])
	root.Right = sortedArrayToBST(nums[mid+1:])
	return root
}
