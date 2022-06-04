package triangle

func MinimumTotal(triangle [][]int) int {
	minlen := triangle[len(triangle)-1]

	for layer := len(triangle) - 2; layer >= 0; layer-- {
		for i := 0; i <= layer; i++ {
			// Find the lesser of its two children, and sum the current value in the triangle with it.
			minlen[i] = min(minlen[i], minlen[i+1]) + triangle[layer][i]
		}
	}

	return minlen[0]
}

func MinimumTotal2(triangle [][]int) int {
	dp := make([]int, len(triangle)) // The min value in i-th level and j-th num
	dp[0] = triangle[0][0]

	for i := 1; i < len(triangle); i++ {
		dp[i] = triangle[i][i] + dp[i-1]

		for j := i - 1; j >= 1; j-- {
			dp[j] = triangle[i][j] + min(dp[j-1], dp[j])
		}

		dp[0] += triangle[i][0]
	}

	min := dp[0]

	for i := 1; i < len(dp); i++ {
		if min > dp[i] {
			min = dp[i]
		}
	}

	return min
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
