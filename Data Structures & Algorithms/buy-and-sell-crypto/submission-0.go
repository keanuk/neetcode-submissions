func maxProfit(prices []int) int {
	profit := 0
	buy := prices[0]
	for i := 0; i < len(prices); i++ {
		sell := prices[i] - buy
		if profit < sell {
			profit = sell
		}
		if prices[i] < buy {
			buy = prices[i]
		}
	}
	return profit
}
