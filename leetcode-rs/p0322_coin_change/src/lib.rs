struct Solution {}

impl Solution {
    pub fn coin_change(coins: Vec<i32>, amount: i32) -> i32 {
        let mut least = std::collections::HashMap::new();
        for val in coins.iter() {
            least.insert(val.clone(), 1);
        }

        Solution::coin_change_recursive(&coins[..], amount, &mut least)
    }

    fn coin_change_recursive(
        coins: &[i32],
        amount: i32,
        memo: &mut std::collections::HashMap<i32, i32>,
    ) -> i32 {
        if let Some(exist) = memo.get(&amount) {
            return exist.clone();
        }

        if amount == 0 {
            return 0;
        }

        let mut min = -1;
        for val in coins {
            if *val > amount {
                continue;
            }

            let mut res = Solution::coin_change_recursive(coins, amount - val, memo);
            if res == -1 {
                continue;
            }

            res += 1;

            if min == -1 || res < min {
                min = res;
            }
        }

        memo.insert(amount, min);
        min
    }
}
