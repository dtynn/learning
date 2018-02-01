package ImplementStrStr

func strStr(haystack string, needle string) int {
	if haystack == needle {
		return 0
	}

	for i := range haystack {
		tail := i + len(needle)
		if tail > len(haystack) {
			break
		}

		if haystack[i:tail] == needle {
			return i
		}
	}

	return -1
}
