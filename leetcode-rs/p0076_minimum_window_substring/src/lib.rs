pub struct Solution {}

use std::collections::HashMap;

impl Solution {
    pub fn min_window(s: String, t: String) -> String {
        let t_bytes = t.into_bytes();
        let t_bytes_len = t_bytes.len();
        if t_bytes_len == 0 {
            return "".to_owned();
        }

        let s_bytes = s.into_bytes();
        let s_bytes_len = s_bytes.len();

        let (bytes_required, mut bytes_count) = t_bytes.iter().fold(
            (HashMap::new(), HashMap::new()),
            |(mut m_required, mut m_count), byte| {
                m_required.entry(*byte).and_modify(|c| *c += 1).or_insert(1);
                m_count.insert(*byte, 0);
                (m_required, m_count)
            },
        );

        let mut byte_pos = vec![];

        for i in 0..s_bytes_len {
            if bytes_required.contains_key(&s_bytes[i]) {
                byte_pos.push(i);
            }
        }

        if byte_pos.len() < t_bytes.len() {
            return "".to_owned();
        }

        let mut min_win = (0, 0);
        let mut left = 0;
        let mut right = 0;

        loop {
            if Solution::is_window(&bytes_count, &bytes_required) {
                let left_pos = byte_pos[left];
                let right_pos = byte_pos[right - 1] + 1;
                let width = right_pos - left_pos;
                let prev_width = min_win.1 - min_win.0;
                if prev_width == 0 || width < prev_width {
                    min_win.0 = left_pos;
                    min_win.1 = right_pos;
                }

                bytes_count
                    .get_mut(&s_bytes[byte_pos[left]])
                    .map(|c| *c -= 1);
                left += 1;
                continue;
            }

            if right == byte_pos.len() {
                break;
            }

            bytes_count
                .get_mut(&s_bytes[byte_pos[right]])
                .map(|c| *c += 1);
            right += 1;
        }

        String::from_utf8(s_bytes[min_win.0..min_win.1].to_vec()).unwrap()
    }

    fn is_window(m_count: &HashMap<u8, usize>, m_required: &HashMap<u8, usize>) -> bool {
        for (b, required) in m_required {
            if required > m_count.get(b).unwrap() {
                return false;
            }
        }

        true
    }
}
