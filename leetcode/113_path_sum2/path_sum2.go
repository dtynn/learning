package path_sum2

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, sum int) [][]int {
	res := make([][]int, 0)
	walk(nil, &res, root, sum)
	return res
}

func walk(pre []int, res *[][]int, root *TreeNode, sum int) {
	if root == nil {
		return
	}

	if root.Left == nil && root.Right == nil {
		if root.Val == sum {
			path := make([]int, len(pre)+1)
			copy(path, pre)
			path[len(pre)] = root.Val

			*res = append(*res, path)
		}

		return
	}

	remain := sum - root.Val
	path := make([]int, len(pre)+1)
	copy(path, pre)
	path[len(pre)] = root.Val

	walk(path, res, root.Left, remain)
	walk(path, res, root.Right, remain)
}
