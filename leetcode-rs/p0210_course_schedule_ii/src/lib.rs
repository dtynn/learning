pub struct Solution {}

use std::collections::{BTreeMap, BTreeSet};

impl Solution {
    pub fn find_order(num_courses: i32, prerequisites: Vec<Vec<i32>>) -> Vec<i32> {
        let mut pres: BTreeMap<i32, Vec<i32>> = BTreeMap::new();
        for p in prerequisites.iter() {
            pres.entry(p[0])
                .and_modify(|l| l.push(p[1]))
                .or_insert_with(|| vec![p[1]]);
        }

        let mut orders = Vec::new();
        let mut visited = BTreeSet::new();

        for c_num in 0..num_courses {
            if !Solution::find(c_num, &pres, &mut orders, &mut visited) {
                return vec![];
            };
        }

        orders
    }

    fn find(
        course_num: i32,
        pres: &BTreeMap<i32, Vec<i32>>,
        orders: &mut Vec<i32>,
        visited: &mut BTreeSet<i32>,
    ) -> bool {
        if visited.contains(&course_num) {
            return false;
        }

        if orders.contains(&course_num) {
            return true;
        }

        visited.insert(course_num);
        let mut can = true;
        match pres.get(&course_num) {
            Some(p) if p.len() > 0 => {
                for pre_num in p {
                    if !Solution::find(*pre_num, pres, orders, visited) {
                        can = false;
                        break;
                    };
                }
            }

            _ => {}
        }

        if can {
            orders.push(course_num);
        }
        visited.remove(&course_num);
        can
    }
}
