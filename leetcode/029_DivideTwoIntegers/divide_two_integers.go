package DivideTwoIntegers

import "math"

func divide(dividend int, divisor int) int {
	negative := 0
	if dividend < 0 {
		negative++
		dividend = -dividend
	}

	if divisor < 0 {
		negative++
		divisor = -divisor
	}

	count := 0
	if divisor == 1 {
		count = dividend
	} else {
		for dividend >= divisor {
			count++
			if count > math.MaxInt32 {
				break
			}

			dividend -= divisor
		}
	}

	if negative == 1 {
		count = -count
	}

	if count >= math.MaxInt32 {
		return math.MaxInt32
	}

	if count <= math.MinInt32 {
		return math.MinInt32
	}

	return count
}
