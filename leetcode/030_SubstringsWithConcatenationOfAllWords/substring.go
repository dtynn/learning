package SubstringsWithConcatenationOfAllWords

func findSubstring(s string, words []string) []int {
	length := len(words[0])
	required := length * len(words)

	res := make([]int, 0, len(s)-required+1)
	copyWords := make([]string, len(words))
	for i := 0; i < len(s)-required+1; i++ {
		copy(copyWords, words)
		if checkConcatenation(s[i:i+required], copyWords, length) {
			res = append(res, i)
		}
	}

	return res
}

func checkConcatenation(s string, words []string, length int) bool {
	check := func(sub string) bool {
		for i := 0; i < len(words); i++ {
			if words[i] == sub {
				words = append(words[:i], words[i+1:]...)
				return true
			}
		}

		return false
	}

	n := len(s) / length
	for i := 0; i < n; i++ {
		sub := s[i*length : i*length+length]
		exists := check(sub)
		if !exists {
			return false
		}
	}

	return len(words) == 0
}
