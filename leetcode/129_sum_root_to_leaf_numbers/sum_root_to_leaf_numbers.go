package sum_root_to_leaf_numbers

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumNumbers(root *TreeNode) int {
	var sum int
	walk(0, root, &sum)
	return sum
}

func walk(pre int, root *TreeNode, sum *int) {
	if root == nil {
		return
	}

	pre = pre*10 + root.Val

	if root.Left == nil && root.Right == nil {
		*sum = *sum + pre
		return
	}

	if root.Left != nil {
		walk(pre, root.Left, sum)
	}

	if root.Right != nil {
		walk(pre, root.Right, sum)
	}
}
