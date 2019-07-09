struct Solution {}

impl Solution {
    pub fn rob(nums: Vec<i32>) -> i32 {
        let size = nums.len();
        if size == 0 {
            return 0;
        }

        let mut amounts = vec![0; size];
        amounts[0] = nums[0];
        let mut max_amount = nums[0];

        for i in 1..size {
            let mut max = nums[i];
            for j in 0..i - 1 {
                max = max.max(nums[i] + amounts[j]);
            }

            max_amount = max_amount.max(max);
            amounts[i] = max;
        }

        max_amount
    }
}
