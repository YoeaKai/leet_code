package longest_substring_without_repeating_characters

func LengthOfLongestSubstring(s string) int {
	max := 0
	flag := false
	result := []byte{}

	for i := 0; i < len(s); i++ {
		flag = false

		for j := 0; j < len(result); j++ {
			if result[j] == s[i] {
				if max < len(result) {
					max = len(result)
				}
				result = append(result[j+1:], s[i])
				flag = true
				break
			}
		}

		if !flag {
			result = append(result, s[i])
		}
	}

	if max < len(result) {
		max = len(result)
	}

	return max
}
