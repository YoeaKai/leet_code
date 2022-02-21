package coin_change_2

// Method 1
// This is the method optimized the space by only using one-dimension array from Method 2.
func Change(amount int, coins []int) int {
	dp := make([]int, amount+1) // The counter of how many combination needed to achieve the sum of coins from 1 to amount.
	dp[0] = 1

	for _, coin := range coins {
		for i := coin; i <= amount; i++ {
			dp[i] += dp[i-coin]
		}
	}

	return dp[amount]
}

// Method 2
func Change2(amount int, coins []int) int {
	dp := make([][]int, len(coins)+1) // The counter records the combination of coins and amount.
	for i := 0; i < len(coins)+1; i++ {
		dp[i] = make([]int, amount+1)
	}
	dp[0][0] = 1

	for i := 1; i < len(coins)+1; i++ {
		dp[i][0] = 1
		for j := 1; j < amount+1; j++ {
			deductCbn := 0 // Combination of amount deducting coin[i]

			if coins[i-1] <= j {
				deductCbn = dp[i][j-coins[i-1]]
			}

			dp[i][j] = dp[i-1][j] + deductCbn // "Not using the ith coin, only using the first i-1 coins to make up the amount." + deductCbn
		}
	}

	return dp[len(coins)][amount]
}
