package SudokuSolver

type cell struct {
	num         int
	unavailable [9]int
}

// 先把所有格子的可能值都记录下来
// 找可能选择最少的格子
// 如果可能选择只有一个, 则填入, 更新其对应行, 列, 块其余格子的可能值, 再沿他的行, 列, 块尝试寻找下一个可以填入的格子
// 如果可能选择有多个, 则分别尝试填入
func solveSudoku(board [][]byte) {
	cells := [9][9]cell{}

}
