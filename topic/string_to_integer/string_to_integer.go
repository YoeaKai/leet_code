package string_to_integer

import (
	"math"
	"strings"
)

func MyAtoi(s string) int {
	pos := 1
	res := 0
	s = strings.TrimSpace(s)
	if len(s) == 0 {
		return res
	}
	i := 0
	if s[i] == '+' {
		i++
		pos = 1
	} else if s[i] == '-' {
		i++
		pos = -1
	}
	for ; i < len(s); i++ {
		if pos*res >= math.MaxInt32 {
			return math.MaxInt32
		}
		if pos*res <= math.MinInt32 {
			return math.MinInt32
		}
		if s[i] < '0' || string(s[i]) > "9" {
			return res * pos
		}
		res = res*10 + int(s[i]) - '0'
	}
	if pos*res >= math.MaxInt32 {
		return math.MaxInt32
	}
	if pos*res <= math.MinInt32 {
		return math.MinInt32
	}
	return pos * res
}
