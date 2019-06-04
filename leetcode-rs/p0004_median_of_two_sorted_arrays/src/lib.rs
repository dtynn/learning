struct Solution {}

impl Solution {
    // median 的统计意义: 集合中一半的数字比它大, 另一半比它小
    // 因此解法的关键是将集合分成等分的两段
    // 认为取 long 的前 n 个和 short 的前 m 个组成小数集合
    // 则满足 long[n] >= short[m-1] && short[m] >= long[n-1]
    // 以 long 数组作基础, 当 n 取值在 [half-short_len..=half] 区间范围内, 检查上述标准是否满足
    pub fn find_median_sorted_arrays(nums1: Vec<i32>, nums2: Vec<i32>) -> f64 {
        let (short, long, short_len, long_len) = if nums1.len() < nums2.len() {
            (nums1.as_slice(), nums2.as_slice(), nums1.len(), nums2.len())
        } else {
            (nums2.as_slice(), nums1.as_slice(), nums2.len(), nums1.len())
        };

        // 较短的长度为 0, 题目中不考虑两个数组同时为 空
        if short_len == 0 {
            return if long_len % 2 == 0 {
                (long[long_len / 2 - 1] + long[long_len / 2]) as f64 / 2f64
            } else {
                long[long_len / 2] as f64
            };
        }

        let half = (short_len + long_len) / 2;
        let mut long_small_part_len = half;

        let (min_of_large, max_of_small) = loop {
            let small_parts = (
                &long[..long_small_part_len],
                &short[..half - long_small_part_len],
            );

            let large_parts = (
                &long[long_small_part_len..],
                &short[half - long_small_part_len..],
            );

            let min_of_large = large_parts
                .0
                .get(0)
                .map(|p| *p)
                .unwrap_or(std::i32::MAX)
                .min(large_parts.1.get(0).map(|p| *p).unwrap_or(std::i32::MAX));

            let max_of_small = small_parts
                .0
                .last()
                .map(|p| *p)
                .unwrap_or(std::i32::MIN)
                .max(small_parts.1.last().map(|p| *p).unwrap_or(std::i32::MIN));

            if min_of_large >= max_of_small {
                break (min_of_large, max_of_small);
            }

            long_small_part_len -= 1;
        };

        if (short_len + long_len) % 2 != 0 {
            min_of_large as f64
        } else {
            (min_of_large + max_of_small) as f64 / 2f64
        }
    }
}

#[cfg(test)]
mod test {
    use super::Solution;

    #[test]
    fn find_median_sorted_arrays_test() {
        let nums1 = vec![1];
        let nums2 = vec![2, 3, 4, 5, 6, 7, 8];

        assert_eq!(Solution::find_median_sorted_arrays(nums1, nums2), 4.5);
    }
}
