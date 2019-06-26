struct Solution {}

impl Solution {
    pub fn merge(nums1: &mut Vec<i32>, m: i32, nums2: &mut Vec<i32>, n: i32) {
        let size_m = m as usize;
        let size_n = n as usize;
        assert!(nums1.len() == size_m + size_n);

        let mut idx_m = 0usize;
        let mut merged_size = size_m;

        for idx_n in 0..size_n {
            let num_from_2 = nums2[idx_n];
            if idx_m == merged_size {
                // 有剩余元素, 按照顺序全部在尾部
                nums1[idx_m] = num_from_2;
                idx_m += 1;
                merged_size += 1;
            } else {
                let mut merged = false;
                while idx_m < merged_size {
                    if num_from_2 < nums1[idx_m] {
                        (&mut nums1[idx_m..].rotate_right(1));
                        nums1[idx_m] = num_from_2;
                        idx_m += 1;
                        merged_size += 1;
                        merged = true;
                        break;
                    } else {
                        idx_m += 1;
                    }
                }

                if !merged {
                    nums1[idx_m] = num_from_2;
                    idx_m += 1;
                    merged_size += 1;
                }
            }
        }
    }
}
