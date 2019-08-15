pub struct Solution {}

impl Solution {
    pub fn title_to_number(s: String) -> i32 {
        let bytes: Vec<u8> = s.bytes().collect();
        bytes
            .into_iter()
            .fold(0, |num, byte| num * 26 + Solution::number(byte))
    }

    fn number(byte: u8) -> i32 {
        (byte - b'A' + 1) as i32
    }
}
