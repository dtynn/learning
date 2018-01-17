package StringToInteger

import (
	"math"
)

func myAtoi(str string) int {
	bytes := []byte(str)

	for len(bytes) > 0 {
		if bytes[0] == ' ' {
			bytes = bytes[1:]
		} else {
			break
		}
	}

	if len(bytes) == 0 {
		return 0
	}

	negative := false

	if bytes[0] == '-' || bytes[0] == '+' {
		negative = bytes[0] == '-'
		bytes = bytes[1:]
		if len(bytes) == 0 {
			return 0
		}
	}

	num := 0
	for i := range bytes {
		if bytes[i] >= '0' && bytes[i] <= '9' {
			num = num*10 + int(bytes[i]-'0')
			if negative && -num < math.MinInt32 {
				return math.MinInt32
			}

			if !negative && num > math.MaxInt32 {
				return math.MaxInt32
			}

		} else {
			break
		}
	}

	if negative && num != 0 {
		num = -num
	}

	return num
}
