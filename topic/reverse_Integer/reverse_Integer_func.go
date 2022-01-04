package reverse_Integer_func

import (
	"strconv"
)

func Reverse(x int) int {
	s := ""

	if x < 0 {
		s += "-"
		x *= -1
	}

	for x != 0 {
		s += strconv.Itoa(x % 10)
		x /= 10
	}

	var ans int
	var err error

	ans, err = strconv.Atoi(s)

	if err != nil || ans > 2147483647 || ans < -2147483648 {
		return 0
	}

	return ans
}
