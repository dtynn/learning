package WildcardMatching

func patternMatchPrefix(bs, prefix []byte) bool {
	if len(bs) < len(prefix) {
		return false
	}

	for i := range prefix {
		if prefix[i] != '?' && prefix[i] != bs[i] {
			return false
		}
	}

	return true
}

func patternMatchSuffix(bs, suffix []byte) bool {
	if len(bs) < len(suffix) {
		return false
	}

	bs = bs[len(bs)-len(suffix):]

	for i := len(suffix) - 1; i >= 0; i-- {
		if suffix[i] != '?' && suffix[i] != bs[i] {
			return false
		}
	}

	return true
}

func search(bs, sub []byte) int {
	for i := 0; i < len(bs)-len(sub)+1; i++ {
		if patternMatchPrefix(bs[i:], sub) {
			return i + len(sub)
		}
	}

	return -1
}

func wildcardMatch(bs []byte, pattern [][]byte) bool {
	// 有剩余未匹配的字节
	for len(bs) > 0 {
		// 没有可尝试的 pattern
		if len(pattern) == 0 {
			return false
		}

		p := pattern[0]
		//	仅 * 匹配, 只可能出现在最后一个
		if len(p) == 1 && p[0] == '*' {
			break
		}

		// 如果是 * + 字节匹配
		if p[0] == '*' {
			// 最后一个 pattern
			if len(pattern) == 1 {
				if !patternMatchSuffix(bs, p[1:]) {
					return false
				}

				bs = bs[len(bs):]
			} else {
				idx := search(bs, p[1:])
				// 无法找到
				if idx == -1 {
					return false
				}

				bs = bs[idx:]
			}

		} else {
			// 必须匹配前缀
			if !patternMatchPrefix(bs, p) {
				return false
			}

			bs = bs[len(p):]
		}

		pattern = pattern[1:]
	}

	for len(pattern) > 0 {
		if len(pattern[0]) != 1 || pattern[0][0] != '*' {
			break
		}

		pattern = pattern[1:]
	}

	return len(pattern) == 0
}

func isMatch(s string, p string) bool {
	if len(p) == 0 {
		return len(s) == 0
	}

	pbs := []byte(p)

	patterns := [][]byte{}
	patterns = append(patterns, []byte{pbs[0]})
	for i := 1; i < len(pbs); i++ {
		b := pbs[i]
		// 当前字节不为 *, 向最后一个 pattern 追加
		if b != '*' {
			patterns[len(patterns)-1] = append(patterns[len(patterns)-1], b)
			continue
		}

		// 重复的 *
		if len(patterns[len(patterns)-1]) == 1 && patterns[len(patterns)-1][0] == '*' {
			continue
		}

		patterns = append(patterns, []byte{b})
	}

	return wildcardMatch([]byte(s), patterns)
}
