package balanced_binary_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	diff := nodeDepth(0, root.Left) - nodeDepth(0, root.Right)

	return diff >= -1 && diff <= 1 && isBalanced(root.Left) && isBalanced(root.Right)
}

func nodeDepth(depth int, root *TreeNode) int {
	if root == nil {
		return depth
	}

	if root.Left == nil && root.Right == nil {
		return depth + 1
	}

	leftDepth := nodeDepth(depth+1, root.Left)
	rightDepth := nodeDepth(depth+1, root.Right)
	if leftDepth > rightDepth {
		return leftDepth
	}

	return rightDepth
}
