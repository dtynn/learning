package ch18

type node struct {
	keys   []int
	childs []*node
}

type btree struct {
	t    int
	root *node
}

func (b *btree) search(k int) (*node, int) {
	return b.searchOnNode(b.root, k)
}

func (b *btree) searchOnNode(n *node, k int) (*node, int) {
	i := 0
	for i < len(n.keys) {
		if n.keys[i] >= k {
			break
		}

		i++
	}

	if i < len(n.keys) && n.keys[i] == k {
		return n, i
	}

	if len(n.childs) == 0 {
		return nil, -1
	}

	return b.searchOnNode(n.childs[i], k)
}

func (b *btree) allocateNode() *node {
	return &node{
		keys:   make([]int, 0, b.t*2-1),
		childs: make([]*node, 0, b.t*2),
	}
}

// 分裂节点 n 的第 i 个子节点
// 如果 n == nil, 代表分裂根节点
func (b *btree) splitChild(n *node, i int) {
	if n == nil {
		oldroot := b.root
		b.root = b.allocateNode()
		b.root.childs = append(b.root.childs, oldroot)
		b.splitChild(b.root, 0)
		return
	}

	targetChild := n.childs[i]
	midKey := targetChild.keys[b.t-1]
	right := b.allocateNode()
	right.keys = append(right.keys, targetChild.keys[b.t:]...)

	targetChild.keys = targetChild.keys[:b.t-1]

	// 非叶子节点
	if len(targetChild.childs) > 0 {
		right.childs = append(right.childs, targetChild.childs[b.t:]...)

		targetChild.childs = targetChild.childs[:b.t]
	}

	newkeys := make([]int, 0, b.t*2-1)
	newkeys = append(newkeys, n.keys[:i]...)
	newkeys = append(newkeys, midKey)
	newkeys = append(newkeys, n.keys[i:]...)

	newchilds := make([]*node, 0, b.t*2)
	newchilds = append(newchilds, n.childs[:i+1]...)
	newchilds = append(newchilds, right)
	newchilds = append(newchilds, n.childs[i+1:]...)

}

// b 树的长高只有通过根节点分裂
func (b *btree) insert(k int) {
	if len(b.root.keys) == b.t*2-1 {
		b.splitChild(nil, -1)
	}

	cur := b.root
	// 自 root 向下, 遇到满节点先分裂
	// 直到找到 k 所处的
	for {
		if len(cur.childs) == 0 {
			break
		}

		childIdx := 0
		for childIdx < len(cur.keys) {
			if k <= cur.keys[childIdx] {
				break
			}
			childIdx++
		}

		// 如果 对应的子节点满, 将之分裂之后重新在当前节点上找对应的子节点
		// 否则向下进入对应的子节点
		if len(cur.childs[childIdx].keys) == b.t*2-1 {
			b.splitChild(cur, childIdx)
		} else {
			cur = cur.childs[childIdx]
		}
	}

	var ki int
	for ki < len(cur.keys) {
		if k <= cur.keys[ki] {
			break
		}
	}

	cur.keys = append(cur.keys, 0)
	copy(cur.keys[ki+1:], cur.keys[ki:len(cur.keys)-1])
	cur.keys[ki] = k
}
