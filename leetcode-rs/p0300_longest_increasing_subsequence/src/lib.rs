struct Solution {}

impl Solution {
    pub fn length_of_lis(nums: Vec<i32>) -> i32 {
        let size = nums.len();
        if size <= 1 {
            return size as i32;
        }

        let mut max = 0;
        let mut lis_memo = vec![];
        for i in 0..size {
            let mut len = 1;
            for j in 0..i {
                if nums[i] > nums[j] {
                    len = len.max(lis_memo[j] + 1);
                }
            }

            max = max.max(len);
            lis_memo.push(len);
        }

        max as i32
    }
}
