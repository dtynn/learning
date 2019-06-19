struct Solution {}

impl Solution {
    pub fn merge(intervals: Vec<Vec<i32>>) -> Vec<Vec<i32>> {
        let size = intervals.len();
        if size == 0 {
            return Vec::new();
        }

        use std::cmp::Ordering;

        let mut intervals = intervals;
        intervals.as_mut_slice().sort_by(|a, b| {
            let first_ord = a[0].cmp(&b[0]);
            if first_ord != Ordering::Equal {
                first_ord
            } else {
                a[1].cmp(&b[1])
            }
        });

        let mut res = Vec::with_capacity(size);
        let mut last = vec![intervals[0][0], intervals[0][1]];

        for i in 1..size {
            let cur = &intervals[i];
            if cur[0] <= last[1] {
                last[1] = cur[1].max(last[1]);
            } else {
                res.push(last.clone());
                last[0] = cur[0];
                last[1] = cur[1];
            }
        }

        res.push(last);

        res
    }
}
