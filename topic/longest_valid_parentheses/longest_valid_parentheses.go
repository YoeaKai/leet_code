package longest_valid_parentheses

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func LongestValidParentheses(s string) int {
	left := 0
	right := 0
	maxlength := 0

	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			left++
		} else {
			right++
		}
		if left == right {
			maxlength = max(maxlength, right*2)
		} else if right >= left {
			left = 0
			right = 0
		}
	}

	left = 0
	right = 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '(' {
			left++
		} else {
			right++
		}
		if left == right {
			maxlength = max(maxlength, left*2)
		} else if left >= right {
			left = 0
			right = 0
		}
	}

	return maxlength
}
