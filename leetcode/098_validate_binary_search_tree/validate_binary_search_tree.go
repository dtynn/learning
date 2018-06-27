package validate_binary_search_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Parent struct {
	Node   *TreeNode
	IsLeft bool
}

func isValidBST(root *TreeNode) bool {
	return isValidWithMemo(root, []*Parent{})
}

func isValidWithMemo(root *TreeNode, parents []*Parent) bool {
	if root == nil {
		return true
	}

	for i := len(parents) - 1; i >= 0; i-- {
		p := parents[i]
		if p.IsLeft {
			if root.Val >= p.Node.Val {
				return false
			}
		} else {
			if root.Val <= p.Node.Val {
				return false
			}
		}
	}

	leftParents := make([]*Parent, len(parents)+1)
	copy(leftParents, parents)
	leftParents[len(parents)] = &Parent{Node: root, IsLeft: true}

	rightParents := make([]*Parent, len(parents)+1)
	copy(rightParents, parents)
	rightParents[len(parents)] = &Parent{Node: root, IsLeft: false}

	return isValidWithMemo(root.Left, leftParents) && isValidWithMemo(root.Right, rightParents)
}
