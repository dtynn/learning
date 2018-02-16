package MultiplyStrings

func multiply(num1 string, num2 string) string {
	if len(num1) == 0 || len(num2) == 0 {
		return ""
	}

	if len(num1) < len(num2) {
		num1, num2 = num2, num1
	}

	b1, b2 := []byte(num1), []byte(num2)
	sum := mul(b1, b2[len(b2)-1])
	for i := len(b2) - 2; i >= 0; i-- {
		tens := uint(len(b2) - i - 1)
		sum = add(sum, ten(mul(b1, b2[i]), tens))
	}

	return string(sum)
}

func ten(base []byte, tens uint) []byte {
	if tens == 0 {
		return base
	}

	zeros := make([]byte, tens)
	for i := 0; i < len(zeros); i++ {
		zeros[i] = '0'
	}

	res := make([]byte, len(base)+len(zeros))
	copy(res[:len(base)], base)
	copy(res[len(base):], zeros)
	return res
}

func mul(base []byte, num byte) []byte {
	if num == '0' {
		return []byte{'0'}
	}

	res := make([]byte, len(base)+1)
	var adv byte

	for i := len(base) - 1; i >= 0; i-- {
		c := (base[i]-'0')*(num-'0') + adv
		if c >= 10 {
			adv = c / 10
			c %= 10
		} else {
			adv = 0
		}

		res[i+1] = c + '0'
	}

	if adv == 0 {
		return res[1:]
	}

	res[0] = adv + '0'
	return res
}

func add(num1, num2 []byte) []byte {
	if len(num1) < len(num2) {
		num1, num2 = num2, num1
	}

	ldiff := len(num1) - len(num2)

	res := make([]byte, len(num1)+1)
	var adv byte

	for i := len(num1) - 1; i >= 0; i-- {
		c1 := num1[i] - '0'
		var c2 byte
		if i-ldiff >= 0 {
			c2 = num2[i-ldiff] - '0'
		}

		sum := c1 + c2 + adv
		if sum >= 10 {
			adv = 1
			sum -= 10
		} else {
			adv = 0
		}

		res[i+1] = sum + '0'
	}

	if adv == 0 {
		return res[1:]
	}

	res[0] = adv + '0'
	return res
}
