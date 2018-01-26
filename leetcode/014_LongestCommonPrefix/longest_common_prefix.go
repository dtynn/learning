package LongestCommonPrefix

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	cprefix := func(s1, s2 []byte) []byte {
		i := 0
		for i < len(s1) && i < len(s2) {
			if s1[i] != s2[i] {
				break
			}

			i++
		}

		return s1[:i]
	}

	prefix := []byte(strs[0])
	for i := 1; i < len(strs); i++ {
		prefix = cprefix(prefix, []byte(strs[i]))
	}

	return string(prefix)
}
