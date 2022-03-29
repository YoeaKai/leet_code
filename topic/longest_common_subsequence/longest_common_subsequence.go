package longest_common_subsequence

func LongestCommonSubsequence(text1 string, text2 string) int {
	var last, tmp int
	dp := make([]int, len(text2))

	for i := 0; i < len(text1); i++ {
		last = dp[0]

		if text1[i] == text2[0] && last != 1 {
			last = 1
		}

		for j := 1; j < len(text2); j++ {
			if text1[i] == text2[j] {
				tmp = dp[j-1] + 1
			} else {
				tmp = max(last, dp[j])
			}
			dp[j-1] = last
			last = tmp
		}

		dp[len(text2)-1] = last
	}

	return dp[len(text2)-1]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
