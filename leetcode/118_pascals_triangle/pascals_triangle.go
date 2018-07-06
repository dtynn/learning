package pascals_triangle

func generate(numRows int) [][]int {
	rows := make([][]int, numRows)

	for i := 0; i < numRows; i++ {
		num := i + 1

		row := make([]int, num)
		rows[i] = row

		row[0], row[num-1] = 1, 1
		if num < 3 {
			continue
		}

		lastRow := rows[i-1]

		for j := 1; j < num-1; j++ {
			row[j] = lastRow[j-1] + lastRow[j]
		}
	}

	return rows
}
