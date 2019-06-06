struct Solution {}

impl Solution {
    pub fn three_sum(nums: Vec<i32>) -> Vec<Vec<i32>> {
        let mut nums_and_counts = std::collections::BTreeMap::new();
        for i in nums {
            nums_and_counts
                .entry(i)
                .and_modify(|count| *count += 1)
                .or_insert(1);
        }

        // keys 是有序的
        let keys: Vec<i32> = nums_and_counts.keys().map(|i_p| *i_p).collect();
        let keys_len = keys.len();
        let mut res = Vec::new();
        let mut exists = std::collections::HashSet::new();

        for i in 0..keys_len {
            // 当 keys[i] > 0 时, 后续的和一定大于 0
            if keys[i] > 0 {
                break;
            }

            let i_remain = nums_and_counts
                .get_mut(&keys[i])
                .map(|count| {
                    *count -= 1;
                    *count
                })
                .unwrap_or(0);

            let start = if i_remain == 0 { i + 1 } else { i };

            exists.clear();

            for j in start..keys_len {
                // 当 keys[i] + keys[j] > 0 时, 后续的和一定大于 0
                if keys[j] > 0 - keys[i] {
                    break;
                }

                let required = 0 - keys[i] - keys[j];
                if exists.contains(&required) {
                    continue;
                }

                nums_and_counts.get_mut(&keys[j]).map(|count| *count -= 1);

                if nums_and_counts
                    .get(&required)
                    .map(|count| *count)
                    .unwrap_or(0)
                    > 0
                {
                    exists.insert(keys[j]);
                    res.push(vec![keys[i], keys[j], required]);
                }

                nums_and_counts.get_mut(&keys[j]).map(|count| *count += 1);
            }

            nums_and_counts.get_mut(&keys[i]).map(|count| *count += 1);
        }

        res
    }
}
