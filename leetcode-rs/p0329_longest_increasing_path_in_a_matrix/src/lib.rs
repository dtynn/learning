pub struct Solution {}

use std::collections::BTreeSet;

impl Solution {
    pub fn longest_increasing_path(matrix: Vec<Vec<i32>>) -> i32 {
        let row_size = matrix.len();
        if row_size == 0 {
            return 0;
        }

        let col_size = matrix[0].len();
        let mut path = vec![];
        let mut visited = BTreeSet::new();
        let mut max_matrix = vec![vec![0; col_size]; row_size];

        let mut max = 0;

        for row in 0..row_size {
            for col in 0..col_size {
                path.clear();
                visited.clear();

                max = max.max(Solution::find_longest_increasing_path(
                    &matrix,
                    (row_size, col_size),
                    (row, col),
                    &mut path,
                    &mut visited,
                    &mut max_matrix,
                ));
            }
        }

        max
    }

    fn find_longest_increasing_path(
        matrix: &[Vec<i32>],
        (row_size, col_size): (usize, usize),
        pos: (usize, usize),
        path: &mut Vec<(usize, usize)>,
        visited: &mut BTreeSet<(usize, usize)>,
        max_matrix: &mut Vec<Vec<i32>>,
    ) -> i32 {
        if visited.contains(&pos) {
            return 0;
        }

        let path_size = path.len();
        if path_size > 0 {
            let last = path[path_size - 1];
            if matrix[last.0][last.1] >= matrix[pos.0][pos.1] {
                return 0;
            }
        }

        let mut max = max_matrix[pos.0][pos.1];
        if max > 0 {
            return max;
        }

        path.push(pos);
        visited.insert(pos);

        let mut next_pos_list = vec![];
        if pos.0 > 0 {
            next_pos_list.push((pos.0 - 1, pos.1));
        }

        if pos.0 < row_size - 1 {
            next_pos_list.push((pos.0 + 1, pos.1));
        }

        if pos.1 > 0 {
            next_pos_list.push((pos.0, pos.1 - 1));;
        }

        if pos.1 < col_size - 1 {
            next_pos_list.push((pos.0, pos.1 + 1));
        }

        for next_pos in next_pos_list.clone() {
            max = max.max(Solution::find_longest_increasing_path(
                matrix,
                (row_size, col_size),
                next_pos,
                path,
                visited,
                max_matrix,
            ));
        }

        // self
        max += 1;

        path.pop();
        visited.remove(&pos);
        max_matrix[pos.0][pos.1] = max;

        max
    }
}

#[cfg(test)]
mod tests {
    #[test]
    fn longest_increasing_path_in_a_matrix_test() {
        let matrix = vec![
            vec![0, 1, 2, 3, 4, 5, 6, 7, 8, 9],
            vec![19, 18, 17, 16, 15, 14, 13, 12, 11, 10],
            vec![20, 21, 22, 23, 24, 25, 26, 27, 28, 29],
            vec![39, 38, 37, 36, 35, 34, 33, 32, 31, 30],
            vec![40, 41, 42, 43, 44, 45, 46, 47, 48, 49],
            vec![59, 58, 57, 56, 55, 54, 53, 52, 51, 50],
            vec![60, 61, 62, 63, 64, 65, 66, 67, 68, 69],
            vec![79, 78, 77, 76, 75, 74, 73, 72, 71, 70],
            vec![80, 81, 82, 83, 84, 85, 86, 87, 88, 89],
            vec![99, 98, 97, 96, 95, 94, 93, 92, 91, 90],
            vec![100, 101, 102, 103, 104, 105, 106, 107, 108, 109],
            vec![119, 118, 117, 116, 115, 114, 113, 112, 111, 110],
            vec![120, 121, 122, 123, 124, 125, 126, 127, 128, 129],
            vec![139, 138, 137, 136, 135, 134, 133, 132, 131, 130],
            vec![0, 0, 0, 0, 0, 0, 0, 0, 0, 0],
        ];

        assert_eq!(super::Solution::longest_increasing_path(matrix), 140);
    }
}
