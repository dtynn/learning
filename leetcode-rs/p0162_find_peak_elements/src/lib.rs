struct Solution {}

impl Solution {
    pub fn find_peak_element(nums: Vec<i32>) -> i32 {
        let size = nums.len();
        assert!(size > 0);
        if size == 1 || nums[0] > nums[1] {
            return 0 as i32;
        }

        if nums[size - 1] > nums[size - 2] {
            return (size - 1) as i32;
        }

        for i in 1..size - 1 {
            if nums[i] > nums[i - 1] && nums[i] > nums[i + 1] {
                return i as i32;
            }
        }

        unreachable!();
    }
}
