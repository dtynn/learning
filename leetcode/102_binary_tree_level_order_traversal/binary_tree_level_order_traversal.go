package binary_tree_level_order_traversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	level := []*TreeNode{root}
	res := [][]int{}

	for {
		if len(level) == 0 {
			break
		}

		vals := make([]int, 0, len(level))
		nextLevel := make([]*TreeNode, 0, 2*len(level))
		for i := range level {
			node := level[i]
			vals = append(vals, node.Val)
			if node.Left != nil {
				nextLevel = append(nextLevel, node.Left)
			}

			if node.Right != nil {
				nextLevel = append(nextLevel, node.Right)
			}
		}

		res = append(res, vals)

		level = nextLevel
	}

	return res
}
