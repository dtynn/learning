struct Solution {}

impl Solution {
    pub fn word_break(s: String, word_dict: Vec<String>) -> bool {
        let bytes = s.as_bytes();
        let mut can_dict =
            word_dict
                .into_iter()
                .fold(std::collections::HashMap::new(), |mut dic, word| {
                    dic.insert(word.into_bytes(), true);
                    dic
                });

        Solution::can_break(&bytes[..], &mut can_dict)
    }

    fn can_break(b: &[u8], dict: &mut std::collections::HashMap<Vec<u8>, bool>) -> bool {
        if b.len() == 0 {
            return true;
        }

        if let Some(can) = dict.get(b).cloned() {
            return can;
        }

        for i in 1..b.len() {
            let left = &b[..i];
            let left_can = Solution::can_break(left, dict);
            dict.insert(left.to_owned(), left_can);
            if !left_can {
                continue;
            }

            let right = &b[i..];
            let right_can = Solution::can_break(right, dict);
            dict.insert(right.to_owned(), right_can);
            if right_can {
                return true;
            }
        }

        dict.insert(b.to_owned(), false);
        false
    }
}
