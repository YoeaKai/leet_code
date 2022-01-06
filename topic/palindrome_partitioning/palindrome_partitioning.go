package palindrome_partitioning

func isPalindrome(s string) bool {
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

func dfs(s string, tmp []string, result *[][]string) {
	for i := 1; i <= len(s); i++ {
		if isPalindrome(s[:i]) {
			if i == len(s) {
				tmpStr := make([]string, len(tmp))
				copy(tmpStr, tmp)
				*result = append(*result, append(tmpStr, s[:i]))
			} else {
				dfs(s[i:], append(tmp, s[:i]), result)
			}
		}
	}
}

func Partition(s string) [][]string {
	var result [][]string
	dfs(s, []string{}, &result)
	return result
}
