struct Solution {}

impl Solution {
    pub fn word_break(s: String, word_dict: Vec<String>) -> Vec<String> {
        let mut word_dict = word_dict;
        word_dict.sort_by_key(|word| word.len());

        let b = s.as_bytes();
        let mut possible = std::collections::HashMap::new();
        for w in word_dict.iter() {
            Solution::do_word_break(w.as_bytes(), &word_dict, &mut possible);
            possible
                .entry(w.clone().into_bytes())
                .and_modify(|strs| strs.push(w.clone()))
                .or_insert(vec![w.clone()]);
        }

        Solution::do_word_break(b, &word_dict, &mut possible);
        possible.remove(b).unwrap_or(vec![])
    }

    fn do_word_break(
        b: &[u8],
        word_dict: &[String],
        possible: &mut std::collections::HashMap<Vec<u8>, Vec<String>>,
    ) {
        let size = b.len();
        if size == 0 {
            return;
        }

        if possible.contains_key(b) {
            return;
        }

        for wstr in word_dict.iter() {
            let wb = wstr.as_bytes();
            let w_size = wb.len();
            if w_size >= size || &b[size - w_size..size] != wb {
                continue;
            }

            let head = &b[..size - w_size];
            Solution::do_word_break(head, word_dict, possible);
            let mut possible_for_wstr: Vec<String> = vec![];
            if let Some(p_for_head) = possible.get(head) {
                for s in p_for_head.iter() {
                    possible_for_wstr.push(vec![s.to_owned(), wstr.to_owned()].join(" "));
                }
            }

            possible
                .entry(b.to_owned())
                .and_modify(|sli| sli.append(&mut possible_for_wstr))
                .or_insert(possible_for_wstr);
        }
    }
}

#[cfg(test)]
mod tests {
    #[test]
    fn word_break_2_test() {
        {
            let s = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaabaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa".to_owned();
            let dict = vec![
                "a".to_owned(),
                "aa".to_owned(),
                "aaa".to_owned(),
                "aaaa".to_owned(),
                "aaaaa".to_owned(),
                "aaaaaa".to_owned(),
                "aaaaaaa".to_owned(),
                "aaaaaaaa".to_owned(),
                "aaaaaaaaa".to_owned(),
                "aaaaaaaaaa".to_owned(),
            ];

            let res = super::Solution::word_break(s, dict);
            println!("{:?}", res);
        }

        {
            let s = "pineapplepenapple".to_owned();
            let dict = vec![
                "apple".to_owned(),
                "pen".to_owned(),
                "applepen".to_owned(),
                "pine".to_owned(),
                "pineapple".to_owned(),
            ];

            let res = super::Solution::word_break(s, dict);
            println!("{:?}", res);
        }
    }
}
