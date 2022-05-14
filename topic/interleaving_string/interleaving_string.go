package interleaving_string

// Method 1, dimensionality reduction from Method 2
func IsInterleave(s1 string, s2 string, s3 string) bool {
	if len(s1)+len(s2) != len(s3) {
		return false
	}

	dp := make([]bool, len(s2)+1)
	dp[0] = true

	for i := 0; i <= len(s1); i++ {
		for j := 0; j <= len(s2); j++ {
			// If s1[i-1] is equal to s3[i+j-1], then dp[j] is unchanged.
			if i != 0 && s1[i-1] != s3[i+j-1] {
				dp[j] = false
			}
			if !dp[j] && j != 0 && s2[j-1] == s3[i+j-1] {
				dp[j] = dp[j-1]
			}
		}
	}

	return dp[len(s2)]
}

// Method 2, dp origin
func IsInterleave2(s1 string, s2 string, s3 string) bool {
	if len(s1)+len(s2) != len(s3) {
		return false
	}

	dp := make([][]bool, len(s1)+1)
	for i := 0; i <= len(s1); i++ {
		dp[i] = make([]bool, len(s2)+1)
	}

	dp[0][0] = true

	// Only take s2 for comparison.
	for j := 0; j < len(s2); j++ {
		if s2[j] == s3[j] {
			dp[0][j+1] = dp[0][j]
		} else {
			break
		}
	}

	for i := 1; i <= len(s1); i++ {
		// Only take s1 for comparison
		if s1[i-1] == s3[i-1] {
			dp[i][0] = dp[i-1][0]
		}

		for j := 1; j <= len(s2); j++ {
			if s1[i-1] == s3[i+j-1] {
				dp[i][j] = dp[i-1][j]
			}
			if !dp[i][j] && s2[j-1] == s3[i+j-1] {
				dp[i][j] = dp[i][j-1]
			}
		}
	}

	return dp[len(s1)][len(s2)]
}
