pub struct Solution {}

use std::collections::{BTreeMap, BTreeSet};

impl Solution {
    pub fn can_finish(num_courses: i32, prerequisites: Vec<Vec<i32>>) -> bool {
        let mut pres: BTreeMap<i32, Vec<i32>> = BTreeMap::new();
        for p in prerequisites.iter() {
            pres.entry(p[0])
                .and_modify(|l| l.push(p[1]))
                .or_insert_with(|| vec![p[1]]);
        }

        let mut cans = BTreeSet::new();
        let mut visited_courses = BTreeSet::new();

        for c_num in 0..num_courses {
            if !Solution::if_can_finish(c_num, &pres, &mut cans, &mut visited_courses) {
                return false;
            }
        }

        true
    }

    fn if_can_finish(
        course_num: i32,
        pres: &BTreeMap<i32, Vec<i32>>,
        cans: &mut BTreeSet<i32>,
        visited_courses: &mut BTreeSet<i32>,
    ) -> bool {
        if visited_courses.contains(&course_num) {
            return false;
        }

        if cans.contains(&course_num) {
            return true;
        }

        visited_courses.insert(course_num);

        let mut can = true;
        if let Some(pres_for_n) = pres.get(&course_num) {
            for pre_num in pres_for_n {
                if !Solution::if_can_finish(*pre_num, pres, cans, visited_courses) {
                    can = false;
                    break;
                };
            }
        }

        if can {
            cans.insert(course_num);
        };

        visited_courses.remove(&course_num);
        return can;
    }
}
