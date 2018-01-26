package LongestSubstringWithoutRepeatingCharacters

func lengthOfLongestSubstring(s string) int {
	var longest int

	exists := map[byte]int{}

	bytes := []byte(s)

	for i, b := range bytes {
		// 出现重复字符
		// 将 exists 重置为上一个重复
		if last, ok := exists[b]; ok {
			exists = map[byte]int{}
			for j := last + 1; j < i; j++ {
				exists[bytes[j]] = j
			}
		}

		exists[b] = i
		if len(exists) > longest {
			longest = len(exists)
		}
	}

	return longest
}
