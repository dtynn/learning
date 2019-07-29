pub struct Solution {}

impl Solution {
    pub fn find_kth_largest(nums: Vec<i32>, k: i32) -> i32 {
        let mut gt = vec![];
        let mut lte = vec![];

        let first = nums[0];
        for i in 1..nums.len() {
            let num = nums[i];
            if num > first {
                gt.push(num);
            } else {
                lte.push(num);
            }
        }

        let gt_len = gt.len() as i32;
        if gt_len == k - 1 {
            return first;
        }

        if gt_len >= k {
            return Solution::find_kth_largest(gt, k);
        }

        Solution::find_kth_largest(lte, k - gt_len - 1)
    }
}
