package minimum_depth_of_binary_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {
	return depth(0, root)
}

func depth(d int, root *TreeNode) int {
	if root == nil {
		return d
	}

	d++

	if root.Left == nil && root.Right == nil {
		return d
	}

	leftDepth := depth(d, root.Left)
	rightDepth := depth(d, root.Right)

	if leftDepth > rightDepth {
		return rightDepth
	}

	return leftDepth
}
