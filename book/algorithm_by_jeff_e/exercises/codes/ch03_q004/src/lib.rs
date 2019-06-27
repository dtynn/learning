pub struct SubA {}

impl SubA {
    // 需要根据切分结果, 去重
    pub fn partitions(s: &str) -> usize {
        let mut can_split_memo: std::collections::HashMap<&[u8], bool> =
            std::collections::HashMap::new();

        let bytes = s.as_bytes();
        let size = bytes.len();
        if size == 0 {
            return 0;
        }

        let mut count = 0;
        if SubA::is_word(bytes) {
            count += 1;
        }

        for i in 1..size {
            let (left, right) = bytes.split_at(i);
            if SubA::can_split(left, &mut can_split_memo)
                && SubA::can_split(right, &mut can_split_memo)
            {
                println!(
                    "split: {:?}, {:?}",
                    String::from_utf8_lossy(left),
                    String::from_utf8_lossy(right)
                );
                count += 1;
            }
        }

        println!("memo:");
        for (k, v) in can_split_memo {
            println!("\t key={}, can={}", String::from_utf8_lossy(k), v);
        }

        count
    }

    fn can_split<'a, 'b: 'a>(
        bytes: &'b [u8],
        memo: &mut std::collections::HashMap<&'a [u8], bool>,
    ) -> bool {
        let size = bytes.len();
        if size == 0 {
            return false;
        }

        if let Some(can) = memo.get(bytes).cloned() {
            return can;
        }

        let is_word = SubA::is_word(bytes);
        memo.insert(&bytes, is_word);
        if is_word {
            return true;
        }

        let mut can = false;
        for i in 1..size {
            let (left, right) = bytes.split_at(i);
            if SubA::can_split(left, memo) && SubA::can_split(right, memo) {
                println!(
                    "check: left={:?}, right={:?}",
                    String::from_utf8_lossy(left),
                    String::from_utf8_lossy(right)
                );
                can = true;
                break;
            }
        }

        memo.insert(&bytes, can);
        can
    }

    fn is_word(s: &[u8]) -> bool {
        match s {
            b"ARTIST" | b"OIL" => true,
            b"ART" | b"IS" | b"TOIL" => true,
            _ => false,
        }
    }
}

#[cfg(test)]
mod tests {
    #[test]
    fn sub_a_test() {
        assert_eq!(super::SubA::partitions("ARTISTOIL"), 2);
    }
}
