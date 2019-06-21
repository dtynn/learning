struct Solution {}

impl Solution {
    // 最简单额方法, 计数排序或者桶排序, 需要两次遍历
    // 一次遍历的思路:
    //  - 从头向尾遍历, 遇到 2 则扔到尾部, 遇到 1 则尝试扔到合适的位置
    pub fn sort_colors(nums: &mut Vec<i32>) {
        let mut count = [0, 0, 0];
        for i in nums.iter() {
            count[*i as usize] += 1;
        }

        let total = nums.len();
        for i in 0..total {
            if count[0] > 0 {
                count[0] -= 1;
                nums[i] = 0;
                continue;
            }

            if count[1] > 0 {
                count[1] -= 1;
                nums[i] = 1;
                continue;
            }

            nums[i] = 2;
        }
    }
}
