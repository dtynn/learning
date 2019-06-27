struct Solution {}

impl Solution {
    pub fn max_profit(prices: Vec<i32>) -> i32 {
        Solution::max_profit_for_slice(&prices[..])
    }

    fn max_profit_for_slice(prices: &[i32]) -> i32 {
        if prices.len() <= 1 {
            return 0;
        }

        let mut max_profit = 0;
        let buy = prices[0];
        for i in 1..prices.len() {
            max_profit = max_profit.max(prices[i] - buy);
        }

        max_profit.max(Solution::max_profit_for_slice(&prices[1..]))
    }
}
