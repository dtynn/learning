pub struct Solution {}

use std::collections::BTreeMap;

impl Solution {
    pub fn get_skyline(buildings: Vec<Vec<i32>>) -> Vec<Vec<i32>> {
        let size = buildings.len();
        if size == 0 {
            return vec![];
        }

        let mut res = vec![];
        let mut temp = BTreeMap::new();
        let mut roofs = BTreeMap::new();

        let mut start_at = 0;
        let mut prev_right_most = buildings[0][1];
        for i in 1..size {
            if buildings[i][0] > prev_right_most {
                Solution::get_skyline_in_block(
                    &buildings[start_at..i],
                    &mut res,
                    &mut temp,
                    &mut roofs,
                );
                start_at = i;
            }
            prev_right_most = prev_right_most.max(buildings[i][1])
        }

        Solution::get_skyline_in_block(&buildings[start_at..size], &mut res, &mut temp, &mut roofs);

        res
    }

    fn get_skyline_in_block(
        buildings: &[Vec<i32>],
        res: &mut Vec<Vec<i32>>,
        temp: &mut BTreeMap<(i32, i32), ()>,
        roofs: &mut BTreeMap<i32, i32>,
    ) {
        let size = buildings.len();
        if size == 0 {
            return;
        }

        if size == 1 {
            res.push(vec![buildings[0][0], buildings[0][2]]);
            res.push(vec![buildings[0][1], 0]);
            return;
        }

        // record roofs
        for b_idx in 0..size {
            let left = buildings[b_idx][0];
            let right = buildings[b_idx][1];
            let h = buildings[b_idx][2];
            let h_idxes = vec![left, left + 1, right - 1, right];
            for idx in h_idxes {
                roofs
                    .entry(idx)
                    .and_modify(|prev| {
                        if *prev < h {
                            *prev = h
                        }
                    })
                    .or_insert(h);
            }

            let zero_idxes = vec![left - 1, right + 1];
            for idx in zero_idxes {
                roofs.entry(idx).or_insert(0);
            }
        }

        for b_idx in 0..size {
            let h = buildings[b_idx][2];
            let range = buildings[b_idx][0]..=buildings[b_idx][1];
            for (_, prev) in roofs.range_mut(range) {
                if *prev < h {
                    *prev = h
                }
            }
        }

        for b_idx in 0..size {
            let left_idx = buildings[b_idx][0];
            let right_idx = buildings[b_idx][1];
            let h = buildings[b_idx][2];

            // 左侧
            let left_top = roofs.get(&left_idx).cloned().unwrap();
            // 左立面是最高的
            if h == left_top {
                // 且左侧都更矮
                let prev = roofs.get(&(left_idx - 1)).cloned().unwrap_or(0);
                if prev < h {
                    temp.insert((left_idx, h), ());
                }
            }

            // 右侧
            let right_top = roofs.get(&right_idx).cloned().unwrap();
            if h == right_top {
                let next = roofs.get(&(right_idx + 1)).cloned().unwrap_or(0);
                if h > next {
                    temp.insert((right_idx, next), ());
                }
            }
        }

        for (point, _) in temp.iter() {
            res.push(vec![point.0, point.1]);
        }

        roofs.clear();
        temp.clear();
    }
}
