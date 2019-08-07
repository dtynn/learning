pub struct Solution {}

use std::collections::HashSet;

impl Solution {
    pub fn num_islands(grid: Vec<Vec<char>>) -> i32 {
        let row_size = grid.len();
        if row_size == 0 {
            return 0;
        }

        let col_size = grid[0].len();

        let mut count = 0;
        let mut visited = HashSet::new();

        for row in 0..row_size {
            for col in 0..col_size {
                let pre_count = visited.len();
                Solution::find_edge_of_island(
                    &grid,
                    (row_size, col_size),
                    (row, col),
                    &mut visited,
                );
                if visited.len() > pre_count {
                    count += 1;
                }
            }
        }

        count
    }

    fn find_edge_of_island(
        grid: &[Vec<char>],
        (row_size, col_size): (usize, usize),
        (row, col): (usize, usize),
        visited: &mut HashSet<(usize, usize)>,
    ) {
        if grid[row][col] == '0' {
            return;
        }

        if visited.insert((row, col)) == false {
            return;
        }

        let mut next_pos_list = vec![];
        if row > 0 {
            next_pos_list.push((row - 1, col));
        }

        if row < row_size - 1 {
            next_pos_list.push((row + 1, col));
        }

        if col > 0 {
            next_pos_list.push((row, col - 1));
        }

        if col < col_size - 1 {
            next_pos_list.push((row, col + 1));
        }

        for next_pos in next_pos_list {
            Solution::find_edge_of_island(grid, (row_size, col_size), next_pos, visited);
        }
    }
}
