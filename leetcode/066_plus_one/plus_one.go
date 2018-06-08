package plus_one

func plusOne(digits []int) []int {
	if len(digits) == 0 {
		return []int{1}
	}

	res := make([]int, len(digits))
	copy(res, digits)

	up := false
	res[len(res)-1] = res[len(res)-1] + 1

	for i := len(res) - 1; i >= 0; i-- {
		if res[i] < 10 {
			break
		}

		res[i] = 0
		if i == 0 {
			up = true
			break
		}

		res[i-1] = res[i-1] + 1
	}

	if up {
		res = append([]int{1}, res...)
	}

	return res
}
