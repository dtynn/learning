struct Solution {}

impl Solution {
    // 对于长度为 n 的输入, 缺失的最小整数要么在 [1, n] 之间, 要么是 n+1;
    // 将所有 非正 数字都变成i32::MAX, 这样所有数字都是正数
    // 然后再次遍历数组, 如果元素不是 i32::MAX 且值在 [1, n] 范围内, 则将 其指向的位置数字置为负数
    // 此时数组内第一个正数即为最小未出现的正数-1
    pub fn first_missing_positive(nums: Vec<i32>) -> i32 {
        let mut nums = nums;
        let max = std::i32::MAX;
        let size = nums.len();
        let mut i = 0usize;
        while i < size {
            if nums[i] <= 0 {
                nums[i] = max;
            }

            i += 1;
        }

        i = 0;
        while i < size {
            let num = nums[i];
            let abs = if num >= 0 { num } else { -num } as usize;
            if abs <= size && nums[abs - 1] > 0 {
                nums[abs - 1] *= -1;
            }

            i += 1;
        }

        i = 0;
        while i < size {
            if nums[i] > 0 {
                break;
            }

            i += 1;
        }

        (i + 1) as i32
    }
}
