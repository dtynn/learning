package n_queens

import (
	"strings"
)

var empty = struct{}{}

func solveNQueens(n int) [][]string {
	solutions := make([][]string, 0)
	e := &exists{
		pos:          make(map[[2]int]struct{}),
		row:          make(map[int]struct{}),
		col:          make(map[int]struct{}),
		crossLeftUp:  make(map[[2]int]struct{}),
		crossRightUp: make(map[[2]int]struct{}),
	}

	solveRow(n, 0, e, &solutions)
	return solutions
}

func solveRow(n, rown int, e *exists, solutions *[][]string) {
	for coln := 0; coln < n; coln++ {
		cur := [2]int{rown, coln}
		cross := crossPos(n, rown, coln)
		_, rowExists := e.row[rown]
		_, colExists := e.col[coln]
		_, crossExists1 := e.crossLeftUp[cross[0]]
		_, crossExists2 := e.crossRightUp[cross[1]]

		if rowExists || colExists || crossExists1 || crossExists2 {
			continue
		}

		e.pos[cur] = empty
		e.row[rown] = empty
		e.col[coln] = empty
		e.crossLeftUp[cross[0]] = empty
		e.crossRightUp[cross[1]] = empty
		if rown == n-1 {
			// 最后一行
			*solutions = append(*solutions, e.toSolution())
		} else {
			solveRow(n, rown+1, e, solutions)
		}

		delete(e.pos, cur)
		delete(e.row, rown)
		delete(e.col, coln)
		delete(e.crossLeftUp, cross[0])
		delete(e.crossRightUp, cross[1])
	}
}

type exists struct {
	pos          map[[2]int]struct{}
	row          map[int]struct{}
	col          map[int]struct{}
	crossLeftUp  map[[2]int]struct{}
	crossRightUp map[[2]int]struct{}
}

func (e *exists) toSolution() []string {
	n := len(e.pos)
	matrix := make([][]string, 0, n)
	for i := 0; i < n; i++ {
		line := make([]string, n)
		for j := 0; j < n; j++ {
			line[j] = "."
		}

		matrix = append(matrix, line)
	}

	for p, _ := range e.pos {
		matrix[p[0]][p[1]] = "Q"
	}

	res := []string{}
	for _, line := range matrix {
		res = append(res, strings.Join(line, ""))
	}
	return res
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
