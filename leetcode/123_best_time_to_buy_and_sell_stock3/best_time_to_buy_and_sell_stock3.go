package best_time_to_buy_and_sell_stock3

func maxProfit(prices []int) int {
	cal := func(prices []int, from, to int) (int, [2]int) {
		interval := [2]int{}

		if to-from <= 1 {
			return 0, interval
		}

		max := 0

		for i := from; i < to-1; i++ {
			for j := i + 1; j < to; j++ {
				profit := prices[j] - prices[i]

				if profit > max {
					max = profit
					interval[0], interval[1] = i, j
				}
			}
		}

		return max, interval
	}

	m, days := cal(prices, 0, len(prices))
	if m <= 0 {
		return 0
	}

	max := m
	if mhead, _ := cal(prices, 0, days[0]); mhead+m > max {
		max = mhead + m
	}

	if mtail, _ := cal(prices, days[1], len(prices)); mtail+m > max {
		max = mtail + m
	}

	for i := days[0] + 1; i < days[1]; i++ {
		m1, _ := cal(prices, days[0], i)
		m2, _ := cal(prices, i, days[1]+1)
		if m1+m2 > max {
			max = m1 + m2
		}
	}

	return max
}
