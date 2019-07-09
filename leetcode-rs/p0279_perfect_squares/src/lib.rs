struct Solution {}

impl Solution {
    pub fn num_squares(n: i32) -> i32 {
        let mut nums = vec![];

        let mut sqr = 1;
        loop {
            let num = sqr * sqr;
            if sqr > n {
                break;
            }

            nums.push(num);
            sqr += 1;
        }

        nums.reverse();

        let mut memo = std::collections::HashMap::new();
        for num in nums.iter() {
            memo.insert(num.clone(), 1);
        }

        Solution::possible_num_squares(n, &nums[..], &mut memo)
    }

    fn possible_num_squares(
        n: i32,
        nums: &[i32],
        memo: &mut std::collections::HashMap<i32, i32>,
    ) -> i32 {
        if let Some(least_opt) = memo.get(&n) {
            return least_opt.clone();
        }

        let mut least = 0;
        for num in nums {
            if n < *num {
                continue;
            }

            let mut possible = Solution::possible_num_squares(n - num, nums, memo);
            if possible == 0 {
                continue;
            }

            possible += 1;

            if least == 0 || possible < least {
                least = possible;
            }
        }

        memo.insert(n, least);
        least
    }
}

#[cfg(test)]
mod tests {
    #[test]
    fn perfect_squares_test() {
        let n = 12;

        assert!(super::Solution::num_squares(n) == 3);
    }
}
