//! Given n non-negative integers a1, a2, ..., an ,where each represents a point at coordinate (i, ai).
//! n vertical lines are drawn such that the two endpoints of line i is at (i, ai) and (i, 0).
//! Find two lines, which together with x-axis forms a container, such that the container contains the most water.
//!
//! Note: You may not slant the container and n is at least 2.
//!

struct Solution {}

impl Solution {
    // left 为左侧高柱的 idx, right 为右侧高柱的 idx
    // 当 height_left < height_right 时, left 的右侧如果出现比 left 更高的高柱, 则有可能比当前容积大
    // 而此时,right 的左侧不可能存在构成大于当前容积的高柱
    // 因此向右移动 left
    //
    // 当 height_left > height_right 时, 情况相反
    //
    // 理论上, 当 height_left == height_right 时, left 和 right 可以同时向对向移动
    pub fn max_area(height: Vec<i32>) -> i32 {
        if height.len() < 2 {
            return 0;
        }

        let (mut left, mut right) = (0, height.len() - 1);
        let mut max = 0i32;

        while left < right {
            let (h_left, h_right) = (height[left], height[right]);
            let wide = (right - left) as i32;

            let h_min;
            if h_left <= h_right {
                h_min = h_left;
                left += 1;
            } else {
                h_min = h_right;
                right -= 1;
            }

            max = max.max(h_min * wide);
        }

        max
    }
}
