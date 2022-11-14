package reverse_words_in_a_string

import "strings"

func ReverseWords(s string) string {
	res := []string{}
	idx := len(s) - 1

	for idx >= 0 {
		for idx >= 0 && s[idx] == ' ' {
			idx--
		}

		if idx < 0 {
			break
		}

		head := idx - 1
		for head >= 0 && s[head] != ' ' {
			head--
		}

		res = append(res, s[head+1:idx+1])
		idx = head
	}

	return strings.Join(res, " ")
}
