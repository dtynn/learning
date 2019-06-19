struct Solution {}

impl Solution {
    pub fn can_jump(nums: Vec<i32>) -> bool {
        let mut max_pos = 0;
        for (i, step) in nums.iter().enumerate() {
            // [0..i] 的最远可达位置小于当前位置
            if max_pos < i {
                return false;
            }

            max_pos = max_pos.max(i + *step as usize);
        }

        max_pos >= nums.len() - 1
    }
}
