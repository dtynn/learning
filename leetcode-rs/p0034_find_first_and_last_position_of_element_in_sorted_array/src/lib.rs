struct Solution {}

impl Solution {
    pub fn search_range(nums: Vec<i32>, target: i32) -> Vec<i32> {
        let size = nums.len();
        let left = Solution::search_left_in_slice(&nums, 0, size, target);
        let right = if left == -1 {
            -1
        } else {
            Solution::search_right_in_slice(&nums, left as usize, size, target)
        };

        vec![left, right]
    }

    fn search_left_in_slice(nums: &[i32], start: usize, end: usize, target: i32) -> i32 {
        if start == end {
            return -1;
        }

        let mid = (start + end) / 2;
        let mid_num = nums[mid];
        if mid_num == target {
            if mid == 0 || nums[mid - 1] < target {
                return mid as i32;
            }

            return Solution::search_left_in_slice(nums, start, mid, target);
        }

        if mid_num < target {
            Solution::search_left_in_slice(nums, mid + 1, end, target)
        } else {
            Solution::search_left_in_slice(nums, start, mid, target)
        }
    }

    fn search_right_in_slice(nums: &[i32], start: usize, end: usize, target: i32) -> i32 {
        if start == end {
            return -1;
        }

        let mid = (start + end) / 2;
        let mid_num = nums[mid];
        if mid_num == target {
            if mid == end - 1 || target < nums[mid + 1] {
                return mid as i32;
            }

            return Solution::search_right_in_slice(nums, mid + 1, end, target);
        }

        if mid_num < target {
            Solution::search_right_in_slice(nums, mid + 1, end, target)
        } else {
            Solution::search_right_in_slice(nums, start, mid, target)
        }
    }
}
