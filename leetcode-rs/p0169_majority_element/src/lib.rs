struct Solution {}

impl Solution {
    pub fn majority_element(nums: Vec<i32>) -> i32 {
        let size = nums.len();
        assert!(size > 0);

        use std::collections::HashMap;
        let mut counts = HashMap::new();
        let mut dedup = vec![];

        for num in nums {
            counts
                .entry(num)
                .and_modify(|e| *e += 1)
                .or_insert_with(|| {
                    dedup.push(num);
                    1
                });
        }

        let is_odd = (size % 2) != 0;
        let required = if is_odd { (size / 2) + 1 } else { size / 2 };

        for num in dedup.iter() {
            if let Some(count) = counts.get(num) {
                if *count >= required {
                    return *num;
                }
            }
        }

        dedup.last().cloned().unwrap()
    }
}
