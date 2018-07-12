package valid_palindrome

func isPalindrome(s string) bool {
	bs := []byte(s)
	res := make([]byte, 0)
	for _, b := range bs {
		if (b >= 'a' && b <= 'z') || (b >= '0' && b <= '9') {
			res = append(res, b)
		}

		if b >= 'A' && b <= 'Z' {
			res = append(res, b-'A'+'a')
		}
	}

	for i := 0; i < len(res)/2; i++ {
		if res[i] != res[len(res)-1-i] {
			return false
		}
	}

	return true
}
