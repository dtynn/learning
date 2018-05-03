package n_queens_2

var empty = struct{}{}

func totalNQueens(n int) int {
	total := 0
	e := &exists{
		row:          make(map[int]struct{}),
		col:          make(map[int]struct{}),
		crossLeftUp:  make(map[[2]int]struct{}),
		crossRightUp: make(map[[2]int]struct{}),
	}

	solveRow(n, 0, e, &total)
	return total
}

func solveRow(n, rown int, e *exists, total *int) {
	for coln := 0; coln < n; coln++ {
		cross := crossPos(n, rown, coln)
		_, rowExists := e.row[rown]
		_, colExists := e.col[coln]
		_, crossExists1 := e.crossLeftUp[cross[0]]
		_, crossExists2 := e.crossRightUp[cross[1]]

		if rowExists || colExists || crossExists1 || crossExists2 {
			continue
		}

		e.row[rown] = empty
		e.col[coln] = empty
		e.crossLeftUp[cross[0]] = empty
		e.crossRightUp[cross[1]] = empty
		if rown == n-1 {
			// 最后一行
			*total = *total + 1
		} else {
			solveRow(n, rown+1, e, total)
		}

		delete(e.row, rown)
		delete(e.col, coln)
		delete(e.crossLeftUp, cross[0])
		delete(e.crossRightUp, cross[1])
	}
}

type exists struct {
	row          map[int]struct{}
	col          map[int]struct{}
	crossLeftUp  map[[2]int]struct{}
	crossRightUp map[[2]int]struct{}
}

func crossPos(n, rown, coln int) [2][2]int {
	res := [2][2]int{}

	r, c := rown, coln
	for r > 0 && c > 0 {
		r--
		c--
	}

	res[0] = [2]int{r, c}

	r, c = rown, coln
	for r > 0 && c < n-1 {
		r--
		c++
	}

	res[1] = [2]int{r, c}
	return res
}
