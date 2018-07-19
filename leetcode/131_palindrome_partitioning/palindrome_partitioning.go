package palindrome_partitioning

func partition(s string) [][]string {
	bs := []byte(s)
	n := len(bs)
	isPalindrome := make([][]bool, n)
	for i := 0; i < n; i++ {
		isPalindrome[i] = make([]bool, n)
	}

	for i := 0; i < n; i++ {
		{
			left := i
			right := i
			for left >= 0 && right < n {
				if bs[left] != bs[right] {
					break
				}

				isPalindrome[left][right] = true

				left--
				right++
			}
		}

		{
			left := i
			right := i + 1
			for left >= 0 && right < n {
				if bs[left] != bs[right] {
					break
				}

				isPalindrome[left][right] = true

				left--
				right++
			}
		}
	}

	res := [][]string{}

	dfs([]string{}, bs, 0, isPalindrome, &res)
	return res
}

func dfs(path []string, bs []byte, start int, isPalindrome [][]bool, res *[][]string) {
	if start == len(bs) {
		if len(path) > 0 {
			*res = append(*res, path)
		}
		return
	}

	for i := start; i < len(bs); i++ {
		if isPalindrome[start][i] {
			nextPath := make([]string, len(path)+1)
			copy(nextPath, path)
			nextPath[len(path)] = string(bs[start : i+1])
			dfs(nextPath, bs, i+1, isPalindrome, res)
		}
	}
}
