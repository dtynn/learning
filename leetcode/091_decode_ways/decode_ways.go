package decode_ways

func numDecodings(s string) int {
	if len(s) == 0 {
		return 0
	}

	bs := []byte(s)
	ways := make([]int, len(bs))

	if isValid(bs[:1]) {
		ways[0] = 1
	}

	if len(bs) > 1 {
		var num1 int
		if isValid(bs[1:2]) {
			num1 += ways[0]
		}

		if isValid(bs[0:2]) {
			num1 += 1
		}

		ways[1] = num1
	}

	for i := 2; i < len(bs); i++ {
		var num int
		if isValid(bs[i : i+1]) {
			num += ways[i-1]
		}

		if isValid(bs[i-1 : i+1]) {
			num += ways[i-2]
		}

		ways[i] = num
	}

	return ways[len(ways)-1]
}

func isValid(digits []byte) bool {
	if len(digits) > 2 {
		return false
	}

	if len(digits) == 1 {
		return digits[0] >= '1' && digits[0] <= '9'
	}

	if digits[0] == '1' && digits[1] >= '0' && digits[1] <= '9' {
		return true
	}

	if digits[0] == '2' && digits[1] >= '0' && digits[1] <= '6' {
		return true
	}

	return false
}
