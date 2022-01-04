package longest_common_prefix

func LongestCommonPrefix(strs []string) string {
	if len(strs) == 0 || strs[0] == "" {
		return ""
	}

	flag := false
	s := strs[0]
	for i := 1; i < len(strs); i++ {
		for j := 0; j < len(strs[i]); j++ {
			if j == len(s) {
				flag = true
				break
			}
			if s[j] != strs[i][j] {
				if j == 0 {
					return ""
				}
				flag = true
				s = s[:j]
				break
			}
		}
		if !flag {
			s = strs[i]
		}
		flag = false
	}
	return s
}
