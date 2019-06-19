struct Solution {}

impl Solution {
    pub fn spiral_order(matrix: Vec<Vec<i32>>) -> Vec<i32> {
        let mut res = Vec::new();
        Solution::spiral_order_append_outter(matrix, 0, &mut res);

        res
    }

    fn spiral_order_append_outter(matrix: Vec<Vec<i32>>, stage: usize, res: &mut Vec<i32>) {
        let row = matrix.len();
        if matrix[stage..row - stage].len() == 0 {
            return;
        }

        let col = matrix[0].len();
        if matrix[0][stage..col - stage].len() == 0 {
            return;
        }

        let width = col - stage * 2;
        if width == 1 {
            for row_n in stage..row - stage {
                res.push(matrix[row_n][stage]);
            }
            return;
        }

        let height = row - stage * 2;
        if height == 1 {
            for col_n in stage..col - stage {
                res.push(matrix[stage][col_n]);
            }
            return;
        }

        let top_range = (0, width - 1);
        let right_range = (top_range.1, top_range.1 + height - 1);
        let bottom_range = (right_range.1, right_range.1 + width - 1);
        let left_range = (bottom_range.1, bottom_range.1 + height - 2);
        let total = left_range.1;

        let mut cordinator = (stage, stage);

        res.push(matrix[cordinator.0][cordinator.1]);
        for step in 0..total {
            match step {
                _n if _n >= top_range.0 && _n < top_range.1 => {
                    cordinator.1 += 1;
                    res.push(matrix[cordinator.0][cordinator.1]);
                }

                _n if _n >= right_range.0 && _n < right_range.1 => {
                    cordinator.0 += 1;
                    res.push(matrix[cordinator.0][cordinator.1]);
                }

                _n if _n >= bottom_range.0 && _n < bottom_range.1 => {
                    cordinator.1 -= 1;
                    res.push(matrix[cordinator.0][cordinator.1]);
                }

                _n if _n >= left_range.0 && _n < left_range.1 => {
                    cordinator.0 -= 1;
                    res.push(matrix[cordinator.0][cordinator.1]);
                }

                _ => {}
            }
        }

        Solution::spiral_order_append_outter(matrix, stage + 1, res);
    }
}
