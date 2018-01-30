package GenerateParentheses

type node struct {
	str         string
	left, right int
	ln, rn      *node
}

func generateParenthesis(n int) []string {
	if n == 0 {
		return []string{}
	}

	root := &node{
		str:   "(",
		left:  n - 1,
		right: n,
	}

	strs := make([]string, 0)
	genNode(root, &strs)
	return strs
}

func genNode(n *node, s *[]string) {
	if n.left == 0 && n.right == 0 {
		*s = append(*s, n.str)
		return
	}

	if n.right > n.left {
		n.rn = &node{
			str:   n.str + ")",
			left:  n.left,
			right: n.right - 1,
		}

		genNode(n.rn, s)
	}

	if n.left > 0 {
		n.ln = &node{
			str:   n.str + "(",
			left:  n.left - 1,
			right: n.right,
		}

		genNode(n.ln, s)
	}
}
