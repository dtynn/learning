struct Solution {}

impl Solution {
    // space problem
    pub fn largest_rectangle_area(heights: Vec<i32>) -> i32 {
        let size = heights.len();
        if size == 0 {
            return 0;
        }

        let mut min_heights = vec![vec![0; size + 1]; size + 1];

        let mut i = 0;
        while i < size {
            let h = heights[i];
            let right_bound = i + 1;

            min_heights[i][right_bound] = h;

            let mut left_bound = i;
            while left_bound > 0 {
                min_heights[left_bound - 1][right_bound] =
                    h.min(min_heights[left_bound - 1][right_bound - 1]);
                left_bound -= 1;
            }

            i += 1;
        }

        let mut max_area = 0;

        for left_bound in 0..size {
            for right_bound in left_bound + 1..size + 1 {
                let area = min_heights[left_bound][right_bound] * (right_bound - left_bound) as i32;
                if area > max_area {
                    max_area = area;
                }
            }
        }

        max_area
    }
}

#[cfg(test)]
mod tests {
    #[test]
    fn largest_rectangle_area_test() {
        let heights = vec![2, 1, 5, 6, 2, 3];
        assert_eq!(super::Solution::largest_rectangle_area(heights), 10);
    }
}
