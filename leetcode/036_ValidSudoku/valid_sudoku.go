package ValidSudoku

func isValidSudoku(board [][]byte) bool {
	rowm := map[byte][]int{}
	colm := map[byte][]int{}
	centerm := map[byte][][2]int{}

	center := func(n int) int {
		switch n {
		case 0, 1, 2:
			return 1

		case 3, 4, 5:
			return 4

		case 6, 7, 8:
			return 7
		}

		return 0
	}

	in := func(nums []int, target int) bool {
		for _, n := range nums {
			if n == target {
				return true
			}
		}

		return false
	}

	centerin := func(centers [][2]int, target [2]int) bool {
		for _, c := range centers {
			if c == target {
				return true
			}
		}

		return false
	}

	for ri, row := range board {
		for ci, cell := range row {
			if cell == '.' {
				continue
			}

			if in(rowm[cell], ri) {
				return false
			}

			if in(colm[cell], ci) {
				return false
			}

			c := [2]int{center(ri), center(ci)}

			if centerin(centerm[cell], c) {
				return false
			}

			rowm[cell] = append(rowm[cell], ri)
			colm[cell] = append(colm[cell], ci)
			centerm[cell] = append(centerm[cell], c)
		}
	}

	return true
}
