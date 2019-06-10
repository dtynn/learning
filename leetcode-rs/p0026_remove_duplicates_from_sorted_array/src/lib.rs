struct Solution {}

// 使用 rotate_left 要反复复制
// 直接将未重复的数值填到 dedeup 结果的尾部
impl Solution {
    pub fn remove_duplicates(nums: &mut Vec<i32>) -> i32 {
        if nums.len() <= 1 {
            return nums.len() as i32;
        }

        let mut dedeup_tail = 1;
        let mut idx = 1;

        while idx < nums.len() {
            if nums[idx] != nums[idx - 1] {
                nums[dedeup_tail] = nums[idx];
                dedeup_tail += 1;
            }

            idx += 1;
        }

        dedeup_tail as i32
    }
}
