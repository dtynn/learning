struct Solution {}

impl Solution {
    // 如果要依次记录每个为 0 的点的坐标, 必然需要 O(m+n) 空间
    // 将值为 0 的点所在的行, 列中非 0 的点都替换为一个正常情况不会出现的值
    // 避免直接替换为 0 的情况下, 右侧/下方的点无法区分
    pub fn set_zeroes(matrix: &mut Vec<Vec<i32>>) {
        let rows = matrix.len();
        if rows == 0 {
            return;
        }

        let cols = matrix[0].len();

        // leetcode use std::i32::MIN in some cases
        let placeholder = std::i32::MIN;

        for row in 0..rows {
            for col in 0..cols {
                if matrix[row][col] != 0 {
                    continue;
                }

                for i in 0..cols {
                    if matrix[row][i] != 0 {
                        matrix[row][i] = placeholder;
                    }
                }

                for j in 0..rows {
                    if matrix[j][col] != 0 {
                        matrix[j][col] = placeholder;
                    }
                }
            }
        }

        for row in 0..rows {
            for col in 0..cols {
                if matrix[row][col] == placeholder {
                    matrix[row][col] = 0;
                }
            }
        }
    }
}
