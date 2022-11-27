package arithmetic_slices_II_subsequence

func NumberOfArithmeticSlices(nums []int) (res int) {
	// dp[i][d] denotes the number of subsequences (have at least 2 numbers)
	// that end with the nums[i] and its common difference is d.
	dp := make([]map[int]int, len(nums))

	for i := 0; i < len(nums); i++ {
		dp[i] = make(map[int]int)

		for j := 0; j < i; j++ {
			diff := nums[i] - nums[j]

			if _, ok := dp[j][diff]; !ok {
				dp[i][diff] += 1
				continue
			}

			//  += dp[j][diff] because it's the number of valid subsequences (have >= 3 elements).
			res += dp[j][diff]
			dp[i][diff] += dp[j][diff] + 1
		}
	}

	return res
}
