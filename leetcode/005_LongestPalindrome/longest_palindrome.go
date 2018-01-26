package LongestPalindrome

func longestPalindrome(s string) string {
	check := func(in []byte) bool {
		for i := 0; i < len(in)/2; i++ {
			if in[i] != in[len(in)-1-i] {
				return false
			}
		}

		return true
	}

	start, end := 0, 1

	bytes := []byte(s)
	lasts := map[byte]int{}

	for i, b := range bytes {
		lastIdx, ok := lasts[b]
		if !ok {
			lasts[b] = i
			continue
		}

		// 有重复的字符
		// 如果前一个字符和当前字符不同, 则更新
		if (bytes[i-1] != b) || (i+1 < len(bytes) && bytes[i+1] != b) {
			lasts[b] = i
		}

		maystart, mayend := lastIdx, i+1
		for maystart >= 0 && mayend <= len(bytes) {
			if !check(bytes[maystart:mayend]) {
				break
			}

			if (mayend - maystart) > (end - start) {
				start, end = maystart, mayend
			}

			maystart--
			mayend++
		}
	}

	return s[start:end]
}
