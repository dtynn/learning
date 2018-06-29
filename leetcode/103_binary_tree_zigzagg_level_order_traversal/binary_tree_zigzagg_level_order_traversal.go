package binary_tree_zigzagg_level_order_traversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	level := []*TreeNode{root}
	levelNum := 0
	res := [][]int{}

	for {
		if len(level) == 0 {
			break
		}

		vals := make([]int, len(level))
		nextLevel := make([]*TreeNode, 0, 2*len(level))
		for i := range level {
			node := level[i]

			if levelNum%2 == 1 {
				i = len(level) - 1 - i
			}

			vals[i] = node.Val
			if node.Left != nil {
				nextLevel = append(nextLevel, node.Left)
			}

			if node.Right != nil {
				nextLevel = append(nextLevel, node.Right)
			}
		}

		res = append(res, vals)

		level = nextLevel
		levelNum++
	}

	return res
}
