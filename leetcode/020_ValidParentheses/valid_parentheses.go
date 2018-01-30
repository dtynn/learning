package ValidParentheses

func isValid(s string) bool {
	pair := map[byte]byte{
		'(': ')',
		'[': ']',
		'{': '}',
	}

	right := map[byte]struct{}{
		')': struct{}{},
		']': struct{}{},
		'}': struct{}{},
	}

	bytes := []byte(s)
	unpaired := make([]byte, 0, len(bytes))

	for i := 0; i < len(bytes); i++ {
		b := bytes[i]
		// 是右侧括号
		if _, ok := right[b]; ok {
			// 没有待匹配的左侧符号
			if len(unpaired) == 0 {
				return false
			}

			// 匹配不成功
			if pair[unpaired[len(unpaired)-1]] != b {
				break
			}

			unpaired = unpaired[:len(unpaired)-1]
			continue
		}

		// 是左侧括号
		if _, ok := pair[b]; ok {
			unpaired = append(unpaired, b)
			continue
		}
	}

	return len(unpaired) == 0
}
