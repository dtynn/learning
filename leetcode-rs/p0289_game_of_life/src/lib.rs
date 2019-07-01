struct Solution {}

impl Solution {
    pub fn game_of_life(board: &mut Vec<Vec<i32>>) {
        let rows = board.len();
        if rows == 0 {
            return;
        }

        let cols = board[0].len();

        for row in 0..rows {
            for col in 0..cols {
                let row_start = row.max(1) - 1;
                let row_end = (row + 2).min(rows);
                let col_start = col.max(1) - 1;
                let col_end = (col + 2).min(cols);

                let mut count_live = 0;
                for neighbor_row in row_start..row_end {
                    for neighbor_col in col_start..col_end {
                        if neighbor_row == row && neighbor_col == col {
                            continue;
                        }

                        if board[neighbor_row][neighbor_col] % 2 != 0 {
                            count_live += 1;
                        }
                    }
                }

                let should_turn = (board[row][col] == 1 && (count_live < 2 || count_live > 3))
                    || (board[row][col] == 0 && count_live == 3);

                if should_turn {
                    board[row][col] += 2;
                }
            }
        }

        for row in 0..rows {
            for col in 0..cols {
                let cur = board[row][col];
                if cur > 1 {
                    board[row][col] = (cur + 1) % 2;
                }
            }
        }
    }
}
