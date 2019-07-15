pub struct Solution {}

impl Solution {
    pub fn single_number(nums: Vec<i32>) -> i32 {
        let mut res = nums[0];

        for i in 1..nums.len() {
            res ^= nums[i];
        }

        res
    }
}
