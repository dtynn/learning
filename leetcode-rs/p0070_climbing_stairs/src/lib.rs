struct Solution {}

impl Solution {
    pub fn climb_stairs(n: i32) -> i32 {
        let mut solutions = vec![0; (n + 1) as usize];
        solutions[0] = 1;
        for i in 1..solutions.len() {
            let mut possible = solutions[i - 1];
            if i >= 2 {
                possible += solutions[i - 2];
            }

            solutions[i] = possible
        }

        solutions.last().cloned().unwrap() as i32
    }
}
