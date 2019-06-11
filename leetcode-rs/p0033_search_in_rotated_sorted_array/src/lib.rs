struct Solution {}

impl Solution {
    pub fn search(nums: Vec<i32>, target: i32) -> i32 {
        let length = nums.len();
        if length == 0 {
            return -1;
        }

        Solution::search_in_slice(&nums, 0, length, target)
    }

    fn search_in_slice(nums: &[i32], start: usize, end: usize, target: i32) -> i32 {
        if start == end {
            return -1;
        }

        let mid = (start + end) / 2;

        if nums[mid] == target {
            return mid as i32;
        }

        // 左侧有元素
        if Solution::may_contain(nums, start, mid, target) {
            return Solution::search_in_slice(nums, start, mid, target);
        }

        // 右侧有元素
        if Solution::may_contain(nums, mid, end, target) {
            return Solution::search_in_slice(nums, mid, end, target);
        }

        -1
    }

    fn may_contain(nums: &[i32], start: usize, end: usize, target: i32) -> bool {
        if start == end {
            return false;
        }

        if end - start == 1 {
            return nums[start] == target;
        }

        let last = end - 1;
        if nums[start] < nums[last] {
            target >= nums[start] && target <= nums[last]
        } else {
            target >= nums[start] || target <= nums[last]
        }
    }
}
