struct Solution {}

impl Solution {
    pub fn exist(board: Vec<Vec<char>>, word: String) -> bool {
        if board.is_empty() {
            return false;
        }

        if word.is_empty() {
            return true;
        }

        let rows = board.len();
        let cols = board[0].len();

        let chars: Vec<char> = word.chars().collect();
        let first = chars[0];

        let mut visited = std::collections::HashSet::new();
        for row in 0..rows {
            for col in 0..cols {
                if board[row][col] == first {
                    visited.clear();
                    visited.insert((row, col));
                    if Solution::find(&board, (row, col), &mut visited, &chars[1..]) {
                        return true;
                    }
                }
            }
        }

        false
    }

    fn find(
        board: &Vec<Vec<char>>,
        point: (usize, usize),
        visited: &mut std::collections::HashSet<(usize, usize)>,
        char_seq: &[char],
    ) -> bool {
        if char_seq.is_empty() {
            return true;
        }

        let next_char = char_seq[0];

        if point.0 > 0 {
            let up = (point.0 - 1, point.1);
            if board[up.0][up.1] == next_char && !visited.contains(&up) {
                visited.insert(up);
                if Solution::find(board, up, visited, &char_seq[1..]) {
                    return true;
                }

                visited.remove(&up);
            }
        }

        if point.0 < board.len() - 1 {
            let down = (point.0 + 1, point.1);
            if board[down.0][down.1] == next_char && !visited.contains(&down) {
                visited.insert(down);
                if Solution::find(board, down, visited, &char_seq[1..]) {
                    return true;
                }

                visited.remove(&down);
            }
        }

        if point.1 > 0 {
            let left = (point.0, point.1 - 1);
            if board[left.0][left.1] == next_char && !visited.contains(&left) {
                visited.insert(left);
                if Solution::find(board, left, visited, &char_seq[1..]) {
                    return true;
                }

                visited.remove(&left);
            }
        }

        if point.1 < board[0].len() - 1 {
            let right = (point.0, point.1 + 1);
            if board[right.0][right.1] == next_char && !visited.contains(&right) {
                visited.insert(right);
                if Solution::find(board, right, visited, &char_seq[1..]) {
                    return true;
                }

                visited.remove(&right);
            }
        }

        false
    }
}
