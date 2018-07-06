package flatten_bianry_tree_to_linked_list

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode) {
	if root == nil {
		return
	}

	if root.Left != nil {
		left := root.Left
		right := root.Right

		leftTail := left
		for leftTail.Right != nil {
			leftTail = leftTail.Right
		}

		root.Left = nil
		root.Right = left
		leftTail.Right = right

	}

	flatten(root.Right)
}
