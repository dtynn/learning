package pascals_triangle2

func getRow(rowIndex int) []int {
	row := make([]int, rowIndex+1)
	for i := 0; i < len(row); i++ {
		row[i] = 1
	}

	if rowIndex <= 1 {
		return row
	}

	for i := 2; i <= rowIndex; i++ {
		for j := i - 1; j > 0; j-- {
			row[j] = row[j] + row[j-1]
		}
	}

	return row
}
