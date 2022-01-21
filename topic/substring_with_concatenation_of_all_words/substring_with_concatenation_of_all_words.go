package substring_with_concatenation_of_all_words

// Method 1
func FindSubstring(s string, words []string) []int {
	ans := []int{}
	wordsMap := map[string]int{}
	n := len(words[0]) * len(words)

	for _, word := range words {
		wordsMap[word]++
	}

	for i := 0; i < len(words[0]); i++ {
		pos := i
		for pos+n <= len(s) {
			exist := map[string]int{}
			times := len(words)

			for ; times > 0; times-- { // Check from the back to the front, it may be skip multiple words at once
				str := s[pos+len(words[0])*(times-1) : pos+len(words[0])*times]
				exist[str]++
				if exist[str] > wordsMap[str] {
					break
				}
			}

			if times == 0 {
				ans = append(ans, pos)
				pos += len(words[0])
			} else {
				pos += len(words[0]) * times
			}

		}
	}
	return ans
}

// Method 2 (Time Limit Exceeded)
func isAns(s string, words []string) bool {
	m := 0
	for l := 0; l < len(words); l++ {
		if s[0] == words[l][0] {
			for m = 1; m < len(words[l]); m++ {
				if s[m] != words[l][m] {
					break
				}
			}
			if m == len(words[l]) {
				if len(words) == 1 {
					return true
				}
				tmp := make([]string, len(words))
				copy(tmp, words)
				if isAns(s[len(words[l]):], append(tmp[:l], words[l+1:]...)) {
					return true
				}
			}
		}
	}
	return false
}

// Status: Time Limit Exceeded
func FindSubstring2(s string, words []string) []int {
	n := len(words) * len(words[0])
	i := 0
	k := 0
	ans := []int{}

	for i = 0; i < len(s)-n+1; i++ {
		for j := 0; j < len(words); j++ {
			if s[i] == words[j][0] {
				for k = 1; k < len(words[j]); k++ {
					if s[i+k] != words[j][k] {
						break
					}
				}
				if k == len(words[j]) {
					if len(words) == 1 {
						ans = append(ans, i)
						break
					}
					tmp := make([]string, len(words))
					copy(tmp, words)
					if isAns(s[i+len(words[j]):i+n], append(tmp[:j], tmp[j+1:]...)) {
						ans = append(ans, i)
						break
					}
				}
			}
		}
	}

	return ans
}
