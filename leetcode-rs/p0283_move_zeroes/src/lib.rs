struct Solution {}

impl Solution {
    pub fn move_zeroes(nums: &mut Vec<i32>) {
        let size = nums.len();
        if size <= 1 {
            return;
        }

        let mut next = 0;
        let mut zeroes = 0;
        for _ in 0..size {
            if nums[next] == 0 {
                (&mut nums[next..size - zeroes]).rotate_left(1);
                zeroes += 1;
            } else {
                next += 1;
            }
        }
    }
}
