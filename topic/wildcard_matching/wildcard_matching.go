package wildcard_matching

func IsMatch(s string, p string) bool {
	var (
		sIdx    int
		pIdx    int
		match   int
		staridx int = -1
	)

	for sIdx < len(s) {
		if pIdx < len(p) && (p[pIdx] == '?' || s[sIdx] == p[pIdx]) {
			sIdx++
			pIdx++
		} else if pIdx < len(p) && p[pIdx] == '*' {
			staridx = pIdx
			match = sIdx
			pIdx++
		} else if staridx != -1 {
			// last p pointer was *
			pIdx = staridx + 1
			match++
			sIdx = match
		} else {
			return false
		}
	}

	//check for remaining characters in pattern
	for pIdx < len(p) && p[pIdx] == '*' {
		pIdx++
	}

	return pIdx == len(p)
}
