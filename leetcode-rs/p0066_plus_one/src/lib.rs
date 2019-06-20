struct Solution {}

impl Solution {
    pub fn plus_one(digits: Vec<i32>) -> Vec<i32> {
        let mut res = vec![0; digits.len() + 1];
        (&mut res[1..]).copy_from_slice(&digits[..]);

        let mut tail = res.len();
        while tail > 0 {
            let idx = tail - 1;
            res[idx] += 1;
            if res[idx] != 10 {
                break;
            }

            res[idx] = 0;
            tail -= 1;
        }

        if res[0] == 0 {
            res[1..].to_owned()
        } else {
            res
        }
    }
}
