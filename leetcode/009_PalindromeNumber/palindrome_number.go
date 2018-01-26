package PalindromeNumber

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	digits := make([]int, 0, 10)
	for x > 0 {
		digits = append(digits, x%10)
		x /= 10
	}

	for i := 0; i < len(digits)/2; i++ {
		if digits[i] != digits[len(digits)-1-i] {
			return false
		}
	}

	return true
}
