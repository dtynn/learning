struct Solution {}

impl Solution {
    pub fn max_profit(prices: Vec<i32>) -> i32 {
        if prices.len() <= 1 {
            return 0;
        }

        let mut max_profit = 0;
        let mut cur = prices[0];
        for sell_idx in 1..prices.len() {
            if prices[sell_idx] > cur {
                max_profit += prices[sell_idx] - cur;
            }

            cur = prices[sell_idx]
        }

        max_profit
    }
}
