package best_time_to_buy_and_sell_stock_II

func MaxProfit(prices []int) (profit int) {
	for i := 1; i < len(prices); i++ {
		if prices[i]-prices[i-1] > 0 {
			profit += prices[i] - prices[i-1]
		}
	}
	return profit
}
