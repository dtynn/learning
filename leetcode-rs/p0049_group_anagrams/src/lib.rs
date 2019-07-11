pub struct Solution {}

impl Solution {
    pub fn group_anagrams(strs: Vec<String>) -> Vec<Vec<String>> {
        let mut m = std::collections::HashMap::new();

        for s in strs {
            let mut b = s.as_bytes().to_vec();
            b.sort();

            m.entry(b)
                .and_modify(|v: &mut Vec<String>| v.push(s.clone()))
                .or_insert_with(|| vec![s]);
        }

        m.drain().map(|(_k, v)| v).collect()
    }
}
