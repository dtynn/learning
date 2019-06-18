struct Solution {}

impl Solution {
    pub fn max_sub_array(nums: Vec<i32>) -> i32 {
        let size = nums.len();
        assert!(size > 0);

        let mut last_max = nums[0];
        let mut max = nums[0];

        for i in 1..size {
            let mut cur_max = nums[i];

            if last_max > 0 {
                cur_max += last_max;
            }

            if cur_max > max {
                max = cur_max;
            }

            last_max = cur_max;
        }

        max
    }
}
