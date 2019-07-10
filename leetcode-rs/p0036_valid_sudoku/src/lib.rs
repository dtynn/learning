struct Solution {}

impl Solution {
    pub fn is_valid_sudoku(board: Vec<Vec<char>>) -> bool {
        let mut nums = [
            [
                std::collections::HashMap::new(),
                std::collections::HashMap::new(),
                std::collections::HashMap::new(),
                std::collections::HashMap::new(),
                std::collections::HashMap::new(),
                std::collections::HashMap::new(),
                std::collections::HashMap::new(),
                std::collections::HashMap::new(),
                std::collections::HashMap::new(),
            ],
            [
                std::collections::HashMap::new(),
                std::collections::HashMap::new(),
                std::collections::HashMap::new(),
                std::collections::HashMap::new(),
                std::collections::HashMap::new(),
                std::collections::HashMap::new(),
                std::collections::HashMap::new(),
                std::collections::HashMap::new(),
                std::collections::HashMap::new(),
            ],
            [
                std::collections::HashMap::new(),
                std::collections::HashMap::new(),
                std::collections::HashMap::new(),
                std::collections::HashMap::new(),
                std::collections::HashMap::new(),
                std::collections::HashMap::new(),
                std::collections::HashMap::new(),
                std::collections::HashMap::new(),
                std::collections::HashMap::new(),
            ],
        ];

        for row_idx in 0..9 {
            for col_idx in 0..9 {
                let c = board[row_idx][col_idx];

                if c == '.' {
                    continue;
                }

                let box_idx = (3 * (row_idx / 3)) + col_idx / 3;
                if nums[0][row_idx].insert(c, ()).is_some()
                    || nums[1][col_idx].insert(c, ()).is_some()
                    || nums[2][box_idx].insert(c, ()).is_some()
                {
                    return false;
                };
            }
        }

        true
    }
}

#[cfg(test)]
mod tests {
    #[test]
    fn valid_sudoku_test() {
        let board = vec![
            vec!['.', '8', '7', '6', '5', '4', '3', '2', '1'],
            vec!['2', '.', '.', '.', '.', '.', '.', '.', '.'],
            vec!['3', '.', '.', '.', '.', '.', '.', '.', '.'],
            vec!['4', '.', '.', '.', '.', '.', '.', '.', '.'],
            vec!['5', '.', '.', '.', '.', '.', '.', '.', '.'],
            vec!['6', '.', '.', '.', '.', '.', '.', '.', '.'],
            vec!['7', '.', '.', '.', '.', '.', '.', '.', '.'],
            vec!['8', '.', '.', '.', '.', '.', '.', '.', '.'],
            vec!['9', '.', '.', '.', '.', '.', '.', '.', '.'],
        ];
    }
}
