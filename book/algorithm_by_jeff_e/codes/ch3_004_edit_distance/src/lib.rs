pub fn edit_distance(a: &str, b: &str) -> usize {
    let a_chars: Vec<char> = a.chars().collect();
    let b_chars: Vec<char> = b.chars().collect();

    let a_len = a_chars.len();
    let b_len = b_chars.len();

    // memo[i][j] 表示 a[0..i] 到 b[0..j] 之间的 edit distance
    let mut memo = vec![vec![0; b_len + 1]; a_len + 1];

    for zero_j in 0..b_len + 1 {
        memo[0][zero_j] = zero_j;
    }

    for i_zero in 0..a_len + 1 {
        memo[i_zero][0] = i_zero;
    }

    for i in 1..a_len + 1 {
        for j in 1..b_len + 1 {
            let replace_ops = if a_chars[i - 1] == b_chars[j - 1] {
                memo[i - 1][j - 1]
            } else {
                memo[i - 1][j - 1] + 1
            };

            let insert_ops = memo[i - 1][j] + 1;
            let delete_ops = memo[i][j - 1] + 1;

            memo[i][j] = replace_ops.min(insert_ops).min(delete_ops);
        }
    }

    memo[a_len][b_len]
}

#[cfg(test)]
mod tests {
    #[test]
    fn edit_distance_test() {
        assert_eq!(super::edit_distance("str", "s"), 2);
    }
}
