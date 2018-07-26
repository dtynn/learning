package word_break

func wordBreak(s string, wordDict []string) bool {

	checkPoints := make([]int, 0, len(s))
	checkPoints = append(checkPoints, -1)

	dict := make(map[string]bool)
	for _, str := range wordDict {
		dict[str] = true
	}

	for i := 0; i < len(s); i++ {
		leng := len(checkPoints)
		for j := 1; j <= leng; j++ {
			tail := s[checkPoints[leng-j]+1 : i+1]
			if dict[tail] {
				checkPoints = append(checkPoints, i)
				break
			}
		}
	}
	return checkPoints[len(checkPoints)-1] == len(s)-1
}
