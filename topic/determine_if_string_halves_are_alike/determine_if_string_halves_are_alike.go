package determine_if_string_halves_are_alike

import "strings"

func HalvesAreAlike(s string) bool {
	var vowels int
	s = strings.ToLower(s)

	for i := 0; i < len(s)/2; i++ {
		if s[i] == 97 || s[i] == 101 || s[i] == 105 || s[i] == 111 || s[i] == 117 {
			vowels++
		}
	}

	for i := len(s) / 2; i < len(s); i++ {
		if s[i] == 97 || s[i] == 101 || s[i] == 105 || s[i] == 111 || s[i] == 117 {
			vowels--
		}
	}

	return vowels == 0
}
