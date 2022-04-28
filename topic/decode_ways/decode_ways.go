package decode_ways

func NumDecodings(s string) int {
	if s[0] == '0' {
		return 0
	}

	pre := 1
	cur := 1

	for i := 1; i < len(s); i++ {
		if s[i-1] == '0' {
			if s[i] == '0' { // Invalid
				return 0
			}
			continue
		}

		if s[i] == '0' {
			if s[i-1] > '2' { // Invalid
				return 0
			}
			cur = pre
		} else if (s[i-1] == '2' && s[i] > '6') || s[i-1] > '2' {
			pre = cur
		} else {
			pre, cur = cur, pre+cur
		}
	}

	return cur
}
