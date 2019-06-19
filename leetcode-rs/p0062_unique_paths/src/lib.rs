struct Solution {}

impl Solution {
    pub fn unique_paths(m: i32, n: i32) -> i32 {
        let less = (m.min(n) - 1) as u64;
        let greater = (m.max(n) - 1) as u64;

        ((greater + 1..=greater + less).fold(1, |prod, e| prod * e)
            / (1..=less).fold(1, |prod, e| prod * e)) as i32
    }
}
