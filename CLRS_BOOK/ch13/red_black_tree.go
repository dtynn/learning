package ch13

const (
	black = true
	red   = false
)

type node struct {
	Key   int
	color bool

	parent, left, right *node
}

func isBlack(n *node) bool {
	return n == nil || n.color
}

type redBlackTree struct {
	root *node
}

func (r *redBlackTree) leftRotate(n *node) {
	right := n.right
	right.parent, n.parent = n.parent, right
	if right.parent == nil {
		r.root = right
	} else if right.parent.left == n {
		right.parent.left = right
	} else {
		right.parent.right = right
	}

	n.right, right.left = right.left, n
	if n.right != nil {
		n.right.parent = n
	}
}

func (r *redBlackTree) rightRotate(n *node) {
	left := n.left
	left.parent, n.parent = n.parent, left
	if left.parent == nil {
		r.root = left
	} else if left.parent.left == n {
		left.parent.left = left
	} else {
		left.parent.right = left
	}

	n.left, left.right = left.right, n
	if n.left != nil {
		n.left.parent = n
	}
}

func (r *redBlackTree) Insert(key int) {
	nw := &node{
		Key:   key,
		color: red,
	}

	node := r.root
	if node == nil {
		r.root = nw
	} else {
		for {
			if key < node.Key {
				if node.left == nil {
					node.left = nw
					nw.parent = node
					break
				}

				node = node.left
			} else {
				if node.right == nil {
					node.right = nw
					nw.parent = node
					break
				}

				node = node.right
			}
		}
	}

	r.insertFixup(nw)
}

func (r *redBlackTree) insertFixup(n *node) {
	// 当 n.parent 是红色时
	for !isBlack(n.parent) {
		// 父节点是祖父节点的左子节点
		nodeIsLeftChild := n.parent.left == n
		var uncle *node
		if nodeIsLeftChild {
			uncle = n.parent.right
		} else {
			uncle = n.parent.left
		}

		// 叔节点是红色
		if !isBlack(uncle) {
			n.parent.parent.color = red
			n.parent.color = black
			uncle.color = black
			n = n.parent.parent
			continue
		}

		// 树节点是
		fatherIsLeftChild := n.parent.parent.left == n.parent.parent
		if fatherIsLeftChild {
			if !nodeIsLeftChild {
				n = n.parent
				r.leftRotate(n)
			}

			n.parent.color = black
			n.parent.parent.color = red
			r.rightRotate(n)
		} else {
			if nodeIsLeftChild {
				n = n.parent
				r.rightRotate(n)
			}

			n.parent.color = black
			n.parent.parent.color = red
			r.leftRotate(n)
		}
	}

	r.root.color = black
}
