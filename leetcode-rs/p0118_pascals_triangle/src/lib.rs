struct Solution {}

impl Solution {
    pub fn generate(num_rows: i32) -> Vec<Vec<i32>> {
        let mut res: Vec<Vec<i32>> = Vec::new();

        for i in 0..num_rows {
            let size = (i + 1) as usize;
            let mut line: Vec<i32> = vec![0; size];
            line[0] = 1;
            line[size - 1] = 1;

            if size > 2 {
                for idx in 1..size - 1 {
                    line[idx] = res[size - 2][idx - 1] + res[size - 2][idx];
                }
            }

            res.push(line);
        }

        res
    }
}
