pub struct Solution {}

use std::collections::BTreeMap;

impl Solution {
    pub fn count_smaller(nums: Vec<i32>) -> Vec<i32> {
        let size = nums.len();
        if size == 0 {
            return vec![];
        }

        let mut counts = vec![0; size];
        let mut num_count = BTreeMap::new();

        for i in 0..size {
            let idx = size - 1 - i;
            let num = nums[idx];

            num_count.entry(num).and_modify(|c| *c += 1).or_insert(1);
            counts[idx] = num_count
                .range(..num)
                .fold(0, |sum, (_, count)| sum + *count);
        }

        counts
    }
}
