package main

func maxProfit(prices []int) int {

	if len(prices) == 0 {
		return 0
	}

	var (
		profit    = 0
		minPrice  = prices[0]
		maxProfit = 0
	)

	for _, currentPrice := range prices {
		if currentPrice < minPrice {
			minPrice = currentPrice
		}

		profit = currentPrice - minPrice
		if profit > maxProfit {
			maxProfit = profit
		}
	}

	return maxProfit
}
