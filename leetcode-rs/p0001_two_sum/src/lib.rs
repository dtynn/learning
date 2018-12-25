#[derive(Debug)]
pub struct Solution {}

impl Solution {
    pub fn two_sum(nums: Vec<i32>, target: i32) -> Vec<i32> {
        let mut first = 0usize;
        while first < nums.len() - 1 {
            let mut second = first + 1;
            while second < nums.len() {
                if nums[first] + nums[second] == target {
                    return vec![first as i32, second as i32];
                }

                second += 1;
            }

            first += 1;
        }

        vec![]
    }
}
