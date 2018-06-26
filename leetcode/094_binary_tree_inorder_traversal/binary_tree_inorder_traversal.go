package binary_tree_inorder_traversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	result := []int{}
	parents := []*TreeNode{}
	cur := root

	for cur != nil || len(parents) > 0 {
		for cur != nil {
			parents = append(parents, cur)
			cur = cur.Left
		}

		parent := parents[len(parents)-1]
		parents = parents[:len(parents)-1]
		result = append(result, parent.Val)

		if parent.Right != nil {
			cur = parent.Right
		}
	}

	return result
}
