package palindrome_number

import (
	"strconv"
)

func IsPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	s := strconv.Itoa(x)
	h := len(s) - 1
	e := 0

	for e < h {
		if s[e] != s[h] {
			return false
		}
		h -= 1
		e += 1
	}

	return true
}
