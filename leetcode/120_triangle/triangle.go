package triangle

func minimumTotal(triangle [][]int) int {
	n := len(triangle)

	if n == 0 {
		return 0
	}

	if n == 1 {
		return triangle[0][0]
	}

	sums := make([]int, n)
	sums[0] = triangle[0][0]

	for row := 1; row < n; row++ {
		for index := row; index >= 0; index-- {
			num := triangle[row][index]
			parents := parent(row, index)
			sum1, sum2 := sums[parents[0]], sums[parents[1]]
			min := sum1 + num
			if sum2 < sum1 {
				min = sum2 + num
			}

			sums[index] = min
		}
	}

	min := sums[0]
	for i := 1; i < len(sums); i++ {
		if sums[i] < min {
			min = sums[i]
		}
	}

	return min
}

func parent(row, index int) [2]int {
	if row == 1 || index == 0 {
		return [2]int{0, 0}
	}

	if row == index {
		return [2]int{row - 1, row - 1}
	}

	return [2]int{index - 1, index}
}
