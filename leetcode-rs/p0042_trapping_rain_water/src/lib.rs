struct Solution {}

impl Solution {
    pub fn trap(height: Vec<i32>) -> i32 {
        let mut left = 0;
        let mut trap = 0;
        while left < height.len() {
            let (right, trapped) = Solution::trap_from(&height, left);
            trap += trapped;
            left = right
        }

        trap
    }

    fn trap_from(heights: &[i32], left: usize) -> (usize, i32) {
        let left_height = heights[left];
        let mut area = 0;
        let mut right = left + 1;
        let mut right_max = left + 1;
        // 如果右侧有高于 left bar 的, 则构成 trap
        // 如果右侧没有高于 left bar 的, 找出其中最高的, 看是否能构成 trap
        while right < heights.len() {
            let right_height = heights[right];
            if right_height >= left_height {
                return (right, area);
            }

            if right_height >= heights[right_max] {
                right_max = right;
            }

            area += left_height - right_height;
            right += 1;
        }

        // 从最高的右侧开始看向左是否能构成 trap
        area = 0;
        let left_most = left;
        let mut left = right_max - 1;
        while left > left_most {
            area += heights[right_max] - heights[left];
            left -= 1;
        }

        (right_max, area)
    }
}
