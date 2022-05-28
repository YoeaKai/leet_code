package target_sum

// Method 1
func FindTargetSumWays(nums []int, target int) int {
	sum := 0

	for _, n := range nums {
		sum += n
	}

	// (target+sum)%2 > 0 has a explain in method 2
	if target < -sum || target > sum || (target+sum)%2 > 0 {
		return 0
	}

	n := len(nums)

	// dp[i][j] means the number of ways for the first i-th element to reach a sum j
	dp := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = make([]int, 2*sum+1)
	}
	dp[0][sum] = 1 // index of dp 0 + sum means sum = 0, 0 means -sum

	for i := 1; i <= n; i++ {
		for j := 0; j < 2*sum+1; j++ {
			if j+nums[i-1] < 2*sum+1 {
				dp[i][j] += dp[i-1][j+nums[i-1]]
			}
			if j-nums[i-1] >= 0 {
				dp[i][j] += dp[i-1][j-nums[i-1]]
			}
		}
	}

	// dp is -sum --> 0 --> +sum
	return dp[n][sum+target]
}

// Method 2
func FindTargetSumWays2(nums []int, target int) int {
	sum := 0

	for _, n := range nums {
		sum += n
	}

	// The original problem statement is equivalent to:
	// Find a subset of nums that need to be positive,
	// and the rest of them is negative, such that the sum is equal to target.

	// sum(P) - sum(N) = target
	// sum(P) + sum(N) + sum(P) - sum(N) = target + sum(P) + sum(N)
	// 2 * sum(P) = target + sum(nums)
	// sum(P) = (target + sum(nums)) / 2

	// The above formula has proved that target + sum(nums) must be even.
	if target < -sum || target > sum || (target+sum)%2 > 0 {
		return 0
	}

	return subsetSum(nums, (target+sum)>>1)
}

func subsetSum(nums []int, t int) int {
	dp := make([]int, t+1)
	dp[0] = 1

	for _, n := range nums {
		for i := t; i >= n; i-- {
			dp[i] += dp[i-n]
		}
	}

	return dp[t]
}
