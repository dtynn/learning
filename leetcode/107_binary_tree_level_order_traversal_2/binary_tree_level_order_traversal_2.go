package binary_tree_level_order_traversal_2

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrderBottom(root *TreeNode) [][]int {
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

	for i := 0; i < len(res)/2; i++ {
		res[i], res[len(res)-1-i] = res[len(res)-1-i], res[i]
	}

	return res
}
