package best_time_to_buy_and_sell_stock

func MaxProfit(prices []int) (res int) {
	min := prices[0]

	for _, p := range prices {
		if p < min {
			min = p
		} else if p-min > res {
			res = p - min
		}
	}

	return res
}
