struct Solution {}

impl Solution {
    pub fn longest_palindrome(s: String) -> String {
        let bytes = s.as_bytes();
        let size = bytes.len();
        let mut longest = (0, 0);

        for i in 0..size {
            let longest_len = longest.1 - longest.0;
            let byte_len = i + 1;
            for start in 0..byte_len - longest_len {
                if Solution::is_palindrome(&bytes[start..byte_len]) {
                    longest.0 = start;
                    longest.1 = byte_len;
                    break;
                }
            }
        }

        String::from_utf8_lossy(&bytes[longest.0..longest.1]).into_owned()
    }

    fn is_palindrome(b: &[u8]) -> bool {
        let size = b.len();
        for i in 0..size / 2 {
            if b[i] != b[size - 1 - i] {
                return false;
            }
        }

        true
    }
}
