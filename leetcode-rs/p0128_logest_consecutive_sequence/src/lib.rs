struct Solution {}

impl Solution {
    pub fn longest_consecutive(nums: Vec<i32>) -> i32 {
        use std::collections::HashMap;
        let mut next_map: HashMap<i32, Option<i32>> = HashMap::new();

        for n in nums {
            next_map.get_mut(&(n - 1)).map(|prev| {
                std::mem::replace(prev, Some(n));
            });

            let next = if next_map.contains_key(&(n + 1)) {
                Some(n + 1)
            } else {
                None
            };

            next_map.insert(n, next);
        }

        let mut longest = 0;
        let mut longest_calculated = HashMap::new();

        for (num, next) in next_map.iter() {
            let mut next_opt = next.as_ref();
            let mut possible_longest = 1;

            while !next_opt.is_none() {
                let next_num = next_opt.unwrap();
                match longest_calculated.get(next_num) {
                    Some(l) => {
                        possible_longest += l;
                        break;
                    }

                    None => {
                        possible_longest += 1;
                        next_opt = next_map.get(next_num).and_then(|inner| inner.as_ref());
                    }
                }
            }

            longest = longest.max(possible_longest);
            longest_calculated.insert(num, possible_longest);
        }

        longest
    }
}
