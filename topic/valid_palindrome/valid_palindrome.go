package valid_palindrome

import "strings"

func IsPalindrome(s string) bool {
	front, back := 0, len(s)-1

	for front < back {
		if !isValidChar(s[front]) {
			front++
			continue
		}

		if !isValidChar(s[back]) {
			back--
			continue
		}

		if !strings.EqualFold(string(s[front]), string(s[back])) {
			return false
		}

		front++
		back--
	}

	return true
}

func isValidChar(ch uint8) bool {
	return (ch >= 48 && ch <= 57) || (ch >= 65 && ch <= 90) || (ch >= 97 && ch <= 122)
}
