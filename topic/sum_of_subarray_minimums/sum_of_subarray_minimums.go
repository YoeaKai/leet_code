package sum_of_subarray_minimums

const Mod = 1000000007 // int(math.Pow(10, 9))+7

// Method 1, Stack
// Refer to: https://leetcode.com/problems/sum-of-subarray-minimums/discuss/2118729/Very-detailed-stack-explanation-O(n)-or-Images-and-comments
func SumSubarrayMins(arr []int) (sum int) {
	arr = append(arr, 0)
	stack := []int{-1}

	for right, val := range arr {
		for len(stack) > 1 && arr[stack[len(stack)-1]] > val {
			index := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			// How much arr[index] needed = Left * Right
			sum += (index - stack[len(stack)-1]) * (right - index) * arr[index]
			sum %= Mod
		}
		stack = append(stack, right)
	}

	return sum
}

// Method 2, Dynamic programming, Time Limit Exceeded.
func SumSubarrayMins2(arr []int) int {
	// dp[start] represent that the minimum value in the array from start to end.
	dp := make([]int, len(arr))
	dp[0] = arr[0]
	sum := dp[0]

	for end := 1; end < len(arr); end++ {
		dp[end] = arr[end]
		sum += dp[end]
		for start := end - 1; start >= 0; start-- {
			dp[start] = min(dp[start], dp[start+1])
			sum += dp[start]
			sum %= Mod
		}
	}

	return sum
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
