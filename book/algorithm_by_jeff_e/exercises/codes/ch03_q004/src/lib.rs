pub struct SubA {}

impl SubA {
    // 需要根据切分结果, 去重
    pub fn partitions(s: &str) -> usize {
        SubA::partitions_for_bytes(s.as_bytes())
    }

    fn partitions_for_bytes(b: &[u8]) -> usize {
        if b.is_empty() {
            return 0;
        }

        let mut count = 0;
        if SubA::is_word(b) {
            count += 1;
        }

        for split_at in 1..b.len() {
            let (left, right) = (&b[0..split_at], &b[split_at..]);
            if SubA::is_word(left) {
                count += SubA::partitions_for_bytes(right);
            }
        }

        count
    }

    fn is_word(s: &[u8]) -> bool {
        match s {
            b"ARTIST" | b"OIL" => true,
            b"ART" | b"IS" | b"TOIL" => true,
            b"TO" | b"IL" => true,
            _ => false,
        }
    }
}

#[cfg(test)]
mod tests {
    #[test]
    fn sub_a_test() {
        assert_eq!(super::SubA::partitions("ARTISTOIL"), 3);
    }
}
