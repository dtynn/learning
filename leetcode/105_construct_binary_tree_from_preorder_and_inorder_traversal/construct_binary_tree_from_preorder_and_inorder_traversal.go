package construct_binary_tree_from_preorder_and_inorder_traversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) != len(inorder) {
		return nil
	}

	if len(preorder) == 0 {
		return nil
	}

	root := &TreeNode{
		Val: preorder[0],
	}

	i := 0
	for i < len(preorder) {
		if inorder[i] == preorder[0] {
			break
		}

		i++
	}

	leftSize := i

	preLeft := preorder[1 : 1+leftSize]
	preRight := preorder[1+leftSize:]

	inLeft := inorder[:i]
	inRight := inorder[i+1:]

	root.Left = buildTree(preLeft, inLeft)
	root.Right = buildTree(preRight, inRight)

	return root
}
