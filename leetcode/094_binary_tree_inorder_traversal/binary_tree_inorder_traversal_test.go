package binary_tree_inorder_traversal

import (
	"testing"
)

func TestTraversal(t *testing.T) {
	t.Run("[1,null,2,3]", func(t *testing.T) {
		root := &TreeNode{
			Val:  1,
			Left: nil,
			Right: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 3,
				},
			},
		}

		t.Log(inorderTraversal(root))
	})
}
