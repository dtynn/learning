package binary_tree_maximum_path_sum

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	if root == nil {
		return 0
	}

	max := root.Val
	searchMax(root, &max)

	return max
}

func maxNum(num ...int) int {
	max := num[0]
	for i := 1; i < len(num); i++ {
		if num[i] > max {
			max = num[i]
		}
	}

	return max
}

func searchMax(root *TreeNode, maxp *int) {
	if root == nil {
		return
	}

	leftMax := walk(root.Val, root.Left)
	rightMax := walk(root.Val, root.Right)
	max := maxNum(root.Val, leftMax, rightMax, leftMax+rightMax-root.Val)
	if max > *maxp {
		*maxp = max
	}

	searchMax(root.Left, maxp)
	searchMax(root.Right, maxp)
}

func walk(sum int, node *TreeNode) int {
	if node == nil {
		return sum
	}

	sum += node.Val
	left := walk(sum, node.Left)
	right := walk(sum, node.Right)
	return maxNum(sum, left, right)
}
