package minimum_window_substring

func MinWindow(s string, t string) string {
	charMap := make(map[byte]int)
	for i := range t {
		charMap[t[i]]++
	}

	counter := len(t)
	if counter > len(s) {
		return ""
	}

	// res can't be s, otherwise it won't be able to judge on line 42.
	res := string(make([]byte, len(s)))
	start, end := 0, 0

	for end < len(s) {
		if v, ok := charMap[s[end]]; ok {
			if v > 0 {
				counter--
			}
			charMap[s[end]]--
		}

		for counter <= 0 {
			if len(res) >= end+1-start {
				res = s[start : end+1]
			}
			if _, ok := charMap[s[start]]; ok {
				charMap[s[start]]++
				if charMap[s[start]] > 0 {
					counter++
				}
			}
			start++
		}

		end++
	}

	if res == string(make([]byte, len(s))) {
		return ""
	}

	return res
}
