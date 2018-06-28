package recover_binary_search_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Res struct {
	Pre *TreeNode
	N1  *TreeNode
	N2  *TreeNode
}

func recoverTree(root *TreeNode) {
	if root == nil {
		return
	}

	var res Res
	solve(root, &res)
	res.N1.Val, res.N2.Val = res.N2.Val, res.N1.Val
}

func solve(root *TreeNode, res *Res) {
	if root == nil {
		return
	}

	solve(root.Left, res)
	if res.Pre == nil {
		res.Pre = root
	}

	if root.Val < res.Pre.Val {
		if res.N1 == nil {
			res.N1 = res.Pre
		}

		res.N2 = root
	}

	res.Pre = root
	solve(root.Right, res)
}
