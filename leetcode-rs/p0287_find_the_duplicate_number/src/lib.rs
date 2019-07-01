struct Solution {}

impl Solution {
    // 1. 指向自身的指针, 例: [3, 1, 3, 4, 2]
    pub fn find_duplicate(nums: Vec<i32>) -> i32 {
        let size = nums.len();
        assert!(size > 0);
        if size <= 2 {
            return nums[0];
        }

        let mut slow = nums[0];
        let mut fast = nums[0];

        loop {
            slow = nums[slow as usize];
            fast = nums[nums[fast as usize] as usize];
            if slow == fast {
                break;
            }
        }

        slow = nums[0];
        while slow != fast {
            slow = nums[slow as usize];
            fast = nums[fast as usize];
        }

        slow
    }
}
