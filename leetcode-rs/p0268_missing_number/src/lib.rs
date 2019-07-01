struct Solution {}

impl Solution {
    pub fn missing_number(nums: Vec<i32>) -> i32 {
        let n = nums.len() as i32;
        assert!(n > 0);
        let count = n + 1;
        let count_is_odd = (count % 2) != 0;
        let mid = count / 2;
        let mut xor = if nums[0] < mid { nums[0] } else { n - nums[0] };

        for num in &nums[1..] {
            xor ^= if *num < mid { *num } else { n - num };
        }

        if count_is_odd {
            xor ^= mid;
        }

        for num in nums {
            if num == xor {
                return n - num;
            }
        }

        xor
    }
}
