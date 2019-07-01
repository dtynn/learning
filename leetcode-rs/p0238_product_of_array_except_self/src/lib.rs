struct Solution {}

impl Solution {
    pub fn product_except_self(nums: Vec<i32>) -> Vec<i32> {
        let mut result = vec![];
        let size = nums.len();
        if size == 0 {
            return result;
        }

        let mut left2right = 1;

        for i in 0..size {
            left2right *= nums[i];
            result.push(left2right);
        }

        if size == 1 {
            return result;
        }

        let mut right2left = 1;
        for i in 0..size - 1 {
            let res_idx = size - i - 1;
            result[res_idx] = result[res_idx - 1] * right2left;
            right2left *= nums[res_idx];
        }

        result[0] = right2left;

        result
    }
}
