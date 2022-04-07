package remove_duplicate_letters

func RemoveDuplicateLetters(s string) string {
	last := make([]int, 26)
	for i := 0; i < len(s); i++ {
		last[s[i]-'a'] = i
	}

	var res []rune
	exist := make([]bool, 26)

	for i, c := range s {
		// Remove the character it will appear later and it's larger than current character.
		for len(res) != 0 && !exist[c-'a'] && res[len(res)-1] > c && last[res[len(res)-1]-'a'] > i {
			exist[res[len(res)-1]-'a'] = false
			res = res[:len(res)-1]
		}

		if !exist[c-'a'] {
			res = append(res, c)
			exist[c-'a'] = true
		}
	}

	return string(res)
}
