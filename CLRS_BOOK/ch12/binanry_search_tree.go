package ch12

type BinarySearchTreeNode struct {
	Key int
	parent,
	left,
	right *BinarySearchTreeNode
}

type BinarySearchTree struct {
	root *BinarySearchTreeNode
}

func (b *BinarySearchTree) Search(key int) *BinarySearchTreeNode {
	x := b.root
	for x != nil {
		if x.Key == key {
			return x
		}

		if x.Key > key {
			x = x.left
		} else {
			x = x.right
		}
	}

	return nil
}

func (b *BinarySearchTree) Minimum() *BinarySearchTreeNode {
	return b.minimum(b.root)
}

func (b *BinarySearchTree) minimum(node *BinarySearchTreeNode) *BinarySearchTreeNode {
	if node == nil {
		return nil
	}

	for node.left != nil {
		node = node.left
	}

	return node
}

func (b *BinarySearchTree) Maximum() *BinarySearchTreeNode {
	return b.maximum(b.root)
}

func (b *BinarySearchTree) maximum(node *BinarySearchTreeNode) *BinarySearchTreeNode {
	if node == nil {
		return nil
	}

	for node.right != nil {
		node = node.right
	}

	return node
}

func (b *BinarySearchTree) Successor(node *BinarySearchTreeNode) *BinarySearchTreeNode {
	if node == nil {
		return nil
	}

	if node.right != nil {
		return b.minimum(node.right)
	}

	for node.parent != nil && node.parent.right == node {
		node = node.parent
	}

	return node.parent
}

func (b *BinarySearchTree) PreDecessor(node *BinarySearchTreeNode) *BinarySearchTreeNode {
	if node == nil {
		return nil
	}

	if node.left != nil {
		return b.maximum(node.left)
	}

	for node.parent != nil && node.parent.left == node {
		node = node.parent
	}

	return node.parent
}

func (b *BinarySearchTree) Insert(key int) {
	newn := &BinarySearchTreeNode{
		Key: key,
	}

	if b.root == nil {
		b.root = newn
		return
	}

	node := b.root
	for {
		if key < node.Key {
			if node.left == nil {
				node.left = newn
				newn.parent = node
				return
			}

			node = node.left

		} else {

			if node.right == nil {
				node.right = newn
				newn.parent = node
				return
			}

			node = node.right
		}
	}
}

func (b *BinarySearchTree) Remove(node *BinarySearchTreeNode) {
	if node == nil {
		return
	}

	switch {
	case node.left == nil && node.right == nil:
		b.transplant(node, nil)

	case node.left == nil && node.right != nil:
		b.transplant(node, node.right)

	case node.right == nil && node.left != nil:
		b.transplant(node, node.left)

	default:
		// 右子树存在, 一定会有一个 successor
		suc := b.minimum(node.right)
		// 非右子节点
		if suc != node.right {
			b.transplant(suc, suc.right)
			suc.right = node.right
			suc.right.parent = suc
		}

		b.transplant(node, suc)
		suc.left = node.left
		suc.left.parent = suc
	}
}

func (b *BinarySearchTree) transplant(oldn, newn *BinarySearchTreeNode) {
	p := oldn.parent

	// root
	if p == nil {
		b.root = newn
	} else if p.left == oldn {
		p.left = newn
	} else {
		p.right = newn
	}

	if newn != nil {
		newn.parent = oldn.parent
	}

}

func InorderTreeWalk(node *BinarySearchTreeNode) []*BinarySearchTreeNode {
	nodes := []*BinarySearchTreeNode{}
	inorderTreeWalk(node, &nodes)

	return nodes
}

func inorderTreeWalk(node *BinarySearchTreeNode, nodes *[]*BinarySearchTreeNode) {
	if node == nil {
		return
	}

	inorderTreeWalk(node.left, nodes)
	*nodes = append(*nodes, node)
	inorderTreeWalk(node.right, nodes)
}
