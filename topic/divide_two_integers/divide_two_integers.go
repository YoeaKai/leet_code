package divide_two_integers

func Divide(dividend int, divisor int) int {
	ans := dividend / divisor

	if ans > 2147483647 {
		return 2147483647
	} else if ans < -2147483648 {
		return -2147483648
	}

	return ans
}
