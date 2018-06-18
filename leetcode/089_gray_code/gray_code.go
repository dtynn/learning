package gray_code

func grayCode(n int) []int {
	if n == 0 {
		return []int{0}
	}

	bits := 1
	result := []int{0, 1}

	for {
		if bits >= n {
			break
		}

		for i := len(result) - 1; i >= 0; i-- {
			result = append(result, (1<<uint(bits))+result[i])
		}
		bits++
	}

	return result
}
