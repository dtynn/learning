package best_time_to_buy_and_sell_stock

func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}

	max := 0
	for i := 0; i < len(prices)-1; i++ {
		for j := i + 1; j < len(prices); j++ {
			if profix := prices[j] - prices[i]; profix > max {
				max = profix
			}
		}
	}

	return max
}
