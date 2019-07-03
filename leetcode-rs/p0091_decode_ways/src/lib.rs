struct Solution {}

impl Solution {
    pub fn num_decodings(s: String) -> i32 {
        let digits = s.as_bytes();
        let mut ways = vec![0; digits.len() + 1];
        ways[0] = 1;
        for prefix_len in 1..=digits.len() {
            let mut possible = if digits[prefix_len - 1] == b'0' {
                0
            } else {
                ways[prefix_len - 1]
            };

            if prefix_len > 1 {
                if Solution::is_valid_2digit(&digits[prefix_len - 2..prefix_len]) {
                    possible += ways[prefix_len - 2];
                }
            }

            ways[prefix_len] = possible
        }

        ways.last().cloned().unwrap()
    }

    fn is_valid_2digit(digit: &[u8]) -> bool {
        match digit[0] {
            b'1' => true,

            b'2' => digit[1] >= b'0' && digit[1] <= b'6',

            _ => false,
        }
    }
}
