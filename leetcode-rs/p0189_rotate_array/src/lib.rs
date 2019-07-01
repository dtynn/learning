struct Solution {}

impl Solution {
    pub fn rotate(nums: &mut Vec<i32>, k: i32) {
        let size = nums.len();
        if size <= 1 {
            return;
        }

        let right_steps = (k as usize) % size;
        if right_steps == 0 {
            return;
        }

        let mut tail = nums[size - 1];

        for _ in 0..right_steps {
            for i in 1..size {
                nums[size - i] = nums[size - i - 1];
            }

            nums[0] = tail;
            tail = nums[size - 1];
        }
    }
}
