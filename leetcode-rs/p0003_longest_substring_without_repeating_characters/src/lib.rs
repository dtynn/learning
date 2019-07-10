struct Solution {}

impl Solution {
    pub fn length_of_longest_substring(s: String) -> i32 {
        let mut longest = 0;
        let bytes = s.as_bytes();

        let mut i = 0;
        let mut exsits = std::collections::HashMap::new();
        while i < bytes.len() {
            if let Some(prev_idx) = exsits.insert(bytes[i], i) {
                longest = longest.max(exsits.len());
                i = prev_idx + 1;
                exsits.clear();
                continue;
            }

            i += 1;
        }

        longest = longest.max(exsits.len());

        longest as i32
    }
}
