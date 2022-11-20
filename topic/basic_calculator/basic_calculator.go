package basic_calculator

func Calculate(s string) (res int) {
	var (
		num   int
		sign  int = 1
		stack []int
	)

	for i := 0; i < len(s); i++ {
		if s[i] >= '0' {
			num = num*10 + int(s[i]-'0')
		} else if s[i] == '+' {
			res += sign * num
			num = 0
			sign = 1
		} else if s[i] == '-' {
			res += sign * num
			num = 0
			sign = -1
		} else if s[i] == '(' {
			stack = append(stack, res)
			stack = append(stack, sign)
			sign = 1
			res = 0
		} else if s[i] == ')' {
			res += sign * num
			res *= stack[len(stack)-1] // the sign before the parenthesis
			res += stack[len(stack)-2] // the res calculated before the parenthesis
			stack = stack[:len(stack)-2]
			num = 0
		}
	}

	return res + sign*num
}
