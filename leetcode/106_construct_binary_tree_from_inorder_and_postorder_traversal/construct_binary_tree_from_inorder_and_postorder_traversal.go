package construct_binary_tree_from_inorder_and_postorder_traversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) != len(postorder) {
		return nil
	}

	if len(inorder) == 0 {
		return nil
	}

	size := len(postorder)

	root := &TreeNode{
		Val: postorder[size-1],
	}

	i := 0
	for i < size {
		if inorder[i] == postorder[size-1] {
			break
		}

		i++
	}
	//
	postLeft := postorder[:i]
	postRight := postorder[i : size-1]

	inLeft := inorder[:i]
	inRight := inorder[i+1:]

	root.Left = buildTree(inLeft, postLeft)
	root.Right = buildTree(inRight, postRight)

	return root
}
