package RegularExpressionMatching

type pattern struct {
	b        byte
	wildcard bool
	next     *pattern
}

func (p *pattern) match(b byte) bool {
	return p.b == '.' || p.b == b
}

func regMatch(bytes []byte, pat *pattern) bool {
	var i int
	for i < len(bytes) {
		b := bytes[i]

		if pat == nil {
			// 有剩余字符, 而没有可以匹配的 pattern
			return false
		}

		// 如果是 *匹配, 先看从下一个 pattern 起能否匹配当前剩余的 bytes
		if pat.wildcard && regMatch(bytes[i:], pat.next) {
			return true
		}

		if !pat.match(b) {
			// 不是 *匹配, 则整个字符串不匹配
			if !pat.wildcard {
				return false
			}

			// 使用 next pattern 尝试匹配当前所有剩余的 bytes
			pat = pat.next
			continue
		}

		// 如果不是 *匹配, 进入下一个 pattern
		if !pat.wildcard {
			pat = pat.next
		}

		// 当前字符匹配成功, 移动到下一个字符
		i++
	}

	// 所有字符已经经过匹配, 此时尝试 skip 掉剩下的所有连续的 *匹配
	for pat != nil {
		if !pat.wildcard {
			break
		}

		pat = pat.next
	}

	// 如果有剩余的非 *匹配, 说明字符串长度不满足
	return pat == nil
}

func isMatch(s string, p string) bool {
	// 生成所有匹配
	pb := []byte(p)
	var first, cur *pattern

	for i, b := range pb {
		if i == 0 {
			first = &pattern{
				b: b,
			}
			cur = first
			continue
		}

		if b == '*' {
			cur.wildcard = true
			continue
		}

		cur.next = &pattern{
			b: b,
		}
		cur = cur.next
	}

	return regMatch([]byte(s), first)
}
