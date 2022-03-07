package different_ways_to_add_parentheses

func DiffWaysToCompute(expression string) []int {
	nums, ops := parse(expression)
	dp := make([][][]int, len(nums))

	// Initialization.
	for i := 0; i < len(nums); i++ {
		dp[i] = make([][]int, len(nums))
		for j := 0; j < len(nums); j++ {
			dp[i][j] = make([]int, 0)
		}
		dp[i][i] = append(dp[i][i], nums[i])
	}

	for l := 2; l <= len(nums); l++ {
		for i := 0; i < len(nums)-l+1; i++ {
			j := i + l - 1
			for k := i; k < j; k++ {
				// Cross calculation to find the all results.
				for _, val1 := range dp[i][k] {
					for _, val2 := range dp[k+1][j] {
						dp[i][j] = append(dp[i][j], calculate(ops[k], val1, val2))
					}
				}
			}
		}
	}

	return dp[0][len(nums)-1]
}

func parse(input string) ([]int, []byte) {
	nums := make([]int, 0)
	ops := make([]byte, 0)
	num := 0

	for i := 0; i < len(input); i++ {
		ch := input[i]
		if ch >= '0' && ch <= '9' {
			num = num*10 + int(ch-'0')
		} else {
			nums = append(nums, num)
			num = 0
			ops = append(ops, ch)
		}
	}

	nums = append(nums, num)
	return nums, ops
}

func calculate(op byte, a, b int) int {
	switch op {
	case '*':
		return a * b
	case '-':
		return a - b
	case '+':
		return a + b
	}
	return 0
}
