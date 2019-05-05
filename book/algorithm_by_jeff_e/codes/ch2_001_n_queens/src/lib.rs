fn place_queens(queens: &mut Vec<usize>, row: usize, solutions: &mut Vec<Vec<usize>>) {
    if queens.len() == row {
        solutions.push(queens.clone());
        return;
    }

    for col in 0..queens.len() {
        let mut legal = true;

        for prev_row in 0..row {
            if queens[prev_row] == col
                || queens[prev_row] == col - row + prev_row
                || queens[prev_row] == col + row - prev_row
            {
                legal = false;
                break;
            }
        }

        if legal {
            queens[row] = col;
            place_queens(queens, row + 1, solutions);
        }
    }
}

#[cfg(test)]
mod test {
    use super::place_queens;

    #[test]
    fn n_queens_test() {
        let n = 8usize;
        let mut empty = vec![0; n];
        let mut solutions = Vec::new();
        place_queens(&mut empty, 0, &mut solutions);
        assert_eq!(solutions.len(), 92);
    }
}
