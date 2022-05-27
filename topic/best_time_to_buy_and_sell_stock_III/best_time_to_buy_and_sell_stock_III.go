package best_time_to_buy_and_sell_stock_III

const MaxUint = ^uint(0)
const MaxInt = int(MaxUint >> 1)
const MinInt = -MaxInt - 1

func MaxProfit(prices []int) int {
	hold1, hold2 := MinInt, MinInt
	release1, release2 := 0, 0

	for _, p := range prices {
		release2 = max(release2, hold2+p) // The maximum if we've just sold 2nd stock so far.
		hold2 = max(hold2, release1-p)    // The maximum if we've just buy  2nd stock so far.
		release1 = max(release1, hold1+p) // The maximum if we've just sold 1nd stock so far.
		hold1 = max(hold1, -p)            // The maximum if we've just buy  1st stock so far.
	}

	return release2 ///Since release1 is initiated as 0, so release2 will always higher than release1.
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
