package scramble_string

func isScramble(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	size := len(s1)

	if s1 == s2 {
		return true
	}

	count := map[byte]int{}
	for i := range s1 {
		count[s1[i]]++
	}

	for i := range s2 {
		count[s2[i]]--
		if count[s2[i]] < 0 {
			return false
		}
	}

	for i := 1; i < size; i++ {
		if isScramble(s1[:i], s2[:i]) && isScramble(s1[i:], s2[i:]) {
			return true
		}

		if isScramble(s1[:i], s2[size-i:]) && isScramble(s1[i:], s2[:size-i]) {
			return true
		}
	}

	return false
}
