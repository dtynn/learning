package palindromePartitioning

func minCut(s string) int {
	bs := []byte(s)
	n := len(bs)

	if n <= 1 {
		return 0
	}

	isPalindrome := make([][]bool, n)
	for i := 0; i < n; i++ {
		isPalindrome[i] = make([]bool, n)
	}

	cuts := make([]int, n)

	for i := 1; i < n; i++ {
		cuts[i] = i
	}

	for right := 0; right < n; right++ {
		for left := right; left >= 0; left-- {
			if bs[left] == bs[right] && (right-left <= 1 || isPalindrome[left+1][right-1]) {
				isPalindrome[left][right] = true

				if left == 0 {
					cuts[right] = 0
				} else if cuts[right] > cuts[left-1]+1 {
					cuts[right] = cuts[left-1] + 1
				}
			}
		}
	}

	return cuts[n-1]
}
