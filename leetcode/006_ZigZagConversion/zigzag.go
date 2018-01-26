package ZigZagConversion

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	size := (numRows - 1) * 2
	rows := make([][]byte, numRows)
	bytes := []byte(s)
	for i, b := range bytes {
		idx := i % size
		if idx >= numRows {
			idx = size - idx
		}

		rows[idx] = append(rows[idx], b)
	}

	res := rows[0]
	for _, row := range rows[1:] {
		res = append(res, row...)
	}

	return string(res)
}
