package best_time_to_buy_and_sell_stock

import "math"

func maxProfit(prices []int) int {
	min := math.MaxInt64
	maxProfit := 0

	for _, price := range prices {
		if price < min {
			min = price
		} else if price-min > maxProfit {
			maxProfit = price - min
		}
	}

	return maxProfit
}
