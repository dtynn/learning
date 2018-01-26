package ReverseInteger

import (
	"math"
)

func reverse(x int) int {
	negative := x < 0
	if negative {
		x = -x
	}

	digits := make([]int, 0, 10)
	for x > 0 {
		digit := x % 10
		x /= 10

		if digit == 0 && len(digits) == 0 {
			continue
		}

		digits = append(digits, digit)

		if len(digits) > 10 {
			return 0
		}
	}

	var res int
	for i := range digits {
		res = res*10 + digits[i]
	}

	if negative {
		res = -res
	}

	if res > math.MaxInt32 || res < math.MinInt32 {
		return 0
	}

	return res
}
