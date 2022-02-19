package length_of_last_word

func LengthOfLastWord(s string) int {
	res := 0
	idx := len(s) - 1

	if s[idx] == ' ' {
		for idx >= 0 && s[idx] == ' ' {
			idx--
		}
	}

	for idx >= 0 && s[idx] != ' ' {
		idx--
		res++
	}

	return res
}
