package distinct_subsequences

// Method 1, dimensionality reduction from Method 2
func NumDistinct(s string, t string) int {
	dp := make([]int, len(s)+1)
	for i := 1; i <= len(s); i++ {
		dp[i] = 1
	}

	before := 1

	for i := 0; i < len(t); i++ {
		for j := 0; j < len(s); j++ {
			tmp := dp[j+1]
			if t[i] == s[j] {
				dp[j+1] = before + dp[j]
			} else {
				dp[j+1] = dp[j]
			}
			before = tmp
		}
		before = 0
	}

	return dp[len(s)]
}

// Method 2, dp origin
func NumDistinct2(s string, t string) int {
	dp := make([][]int, len(t)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(s)+1)
	}

	// Filling the first row.
	for i := 0; i <= len(s); i++ {
		dp[0][i] = 1
	}

	// The first column is 0 by default in every other rows but the first.
	for i := 0; i < len(t); i++ {
		for j := 0; j < len(s); j++ {
			if t[i] == s[j] {
				// this times = the times before t[i+1] and s[j+1] + front latters of s times
				dp[i+1][j+1] = dp[i][j] + dp[i+1][j]
			} else {
				dp[i+1][j+1] = dp[i+1][j]
			}
		}
	}

	return dp[len(t)][len(s)]
}
