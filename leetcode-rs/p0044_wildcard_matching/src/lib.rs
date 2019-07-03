struct Solution {}

impl Solution {
    pub fn is_match(s: String, p: String) -> bool {
        let s_bytes = s.as_bytes();
        let p_bytes = Solution::compact_pattern(p);
        let mut memo = vec![vec![None; p_bytes.len() + 1]; s_bytes.len() + 1];
        memo[s_bytes.len()][p_bytes.len()] = Some(true);
        let res = Solution::is_match_slice(&s_bytes[..], 0, &p_bytes[..], 0, &mut memo);

        res
    }

    fn compact_pattern(p: String) -> Vec<u8> {
        let mut bytes = vec![];

        let raw = p.as_bytes();
        if raw.is_empty() {
            return bytes;
        }

        let mut last = raw[0];
        for i in 1..raw.len() {
            let b = raw[i];
            if b == b'*' && last == b'*' {
                continue;
            }

            bytes.push(last);
            last = b;
        }

        bytes.push(last);
        bytes
    }

    fn is_match_slice(
        b: &[u8],
        b_start: usize,
        p: &[u8],
        p_start: usize,
        memo: &mut Vec<Vec<Option<bool>>>,
    ) -> bool {
        // 每个 element 代表的含义:  memo[b_start][p_start] 代表 b[b_start..] 与 p[p_start..]
        // 是否匹配
        if let Some(res) = memo[b_start][p_start] {
            return res;
        }

        if b.len() - b_start == 0 {
            // both empty
            if p.len() - p_start == 0 {
                memo[b_start][p_start] = Some(true);
                return true;
            }

            // multiple "*"
            if p[p_start] == b'*' {
                let is_match = Solution::is_match_slice(b, b_start, p, p_start + 1, memo);
                memo[b_start][p_start] = Some(is_match);
                return is_match;
            }

            memo[b_start][p_start] = Some(false);
            return false;
        }

        // no pattern left
        if p.len() - p_start == 0 {
            memo[b_start][p_start] = Some(false);
            return false;
        }

        let is_match = match p[p_start] {
            b'*' => {
                // * matches multiple chars
                Solution::is_match_slice(b, b_start+1, p, p_start, memo)
                // * matches only current char
                    || Solution::is_match_slice(b, b_start+1, p, p_start+1, memo)
                // * matches no char at all
                    || Solution::is_match_slice(b, b_start, p, p_start+1, memo)
            }

            _ => {
                (p[p_start] == b'?' || b[b_start] == p[p_start])
                    && Solution::is_match_slice(b, b_start + 1, p, p_start + 1, memo)
            }
        };

        memo[b_start][p_start] = Some(is_match);
        is_match
    }
}

#[cfg(test)]
mod tests {
    #[test]
    fn wildcard_matching_test() {
        let s = "babaaababaabababbbbbbaabaabbabababbaababbaaabbbaaab".to_owned();
        let p = "***bba**a*bbba**aab**b".to_owned();

        assert_eq!(super::Solution::is_match(s, p), false);
    }
}
