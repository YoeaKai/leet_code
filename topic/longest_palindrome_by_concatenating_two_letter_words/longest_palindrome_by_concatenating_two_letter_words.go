package longest_palindrome_by_concatenating_two_letter_words

// Method 1: Slice with 26 * 26 int, Runtime: 401 ms, Memory Usage: 17.1 MB.
func LongestPalindrome(words []string) (res int) {
	m := make([][]int, 26)
	for i := 0; i < 26; i++ {
		m[i] = make([]int, 26)
	}

	for _, word := range words {
		if count := m[word[1]-'a'][word[0]-'a']; count == 0 {
			m[word[0]-'a'][word[1]-'a']++
		} else {
			m[word[1]-'a'][word[0]-'a']--
			res += 4
		}
	}

	for i := 0; i < 26; i++ {
		if m[i][i] > 0 {
			res += 2
			break
		}
	}

	return res
}

// Method 2: Hash map with string keys, Runtime: 469 ms, Memory Usage: 13.4 MB.
func LongestPalindrome2(words []string) (res int) {
	m := make(map[string]int)

	for idx, word := range words {
		if count := m[words[idx][1:]+words[idx][:1]]; count <= 0 {
			m[word]++
		} else {
			m[words[idx][1:]+words[idx][:1]]--
			res += 4
		}
	}

	for key := range m {
		if key[0] == key[1] && m[key] > 0 {
			res += 2
			break
		}
	}

	return res
}
