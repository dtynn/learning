struct Solution {}

impl Solution {
    // 1. 在  .* 时, 持续匹配
    //    例: "ab", ".*""
    // 2. 有* 的子模式被匹配到, 但是仍然有可能需要排除当前模式
    //    例: "bbbbbbba", ".*a*a"
    pub fn is_match(s: String, p: String) -> bool {
        let bytes = s.as_bytes();
        let components = Solution::pattern_components(p);
        let components_ref: Vec<&[u8]> = components.iter().map(|v| v.as_ref()).collect();

        Solution::match_components(&bytes[..], &components_ref[..])
    }

    fn pattern_components(p: String) -> Vec<Vec<u8>> {
        let bytes = p.as_bytes();
        let size = bytes.len();
        if size == 0 {
            return vec![];
        }

        let mut components = vec![];
        let mut cur = vec![bytes[0]];
        for i in 1..size {
            let b = bytes[i];
            match b {
                b'*' => cur.push(b),
                _ => {
                    components.push(cur.clone());
                    cur.clear();
                    cur.push(b);
                }
            }
        }

        components.push(cur.clone());
        components
    }

    fn match_components(b: &[u8], components: &[&[u8]]) -> bool {
        // 必须 1:1 匹配的组件
        let must_match_components = components.iter().fold(0, |mut count, component| {
            if component.len() == 1 {
                count += 1;
            }

            count
        });

        if b.len() == 0 {
            return must_match_components == 0;
        }

        if components.len() == 0 || b.len() < must_match_components {
            return false;
        }

        let first = &components[0];
        if first.len() == 1 {
            if first[0] == b'.' || first[0] == b[0] {
                return Solution::match_components(&b[1..], &components[1..]);
            }

            return false;
        }

        if first[0] == b'.' || first[0] == b[0] {
            if Solution::match_components(&b[1..], components)
                || Solution::match_components(&b[1..], &components[1..])
            {
                return true;
            };
        }

        Solution::match_components(b, &components[1..])
    }
}

#[cfg(test)]
mod tests {
    #[test]
    fn regular_expression_match_test() {
        let s = "bbbbbbba".to_owned();
        let p = ".*a*a".to_owned();

        assert_eq!(super::Solution::is_match(s, p), true);
    }
}
