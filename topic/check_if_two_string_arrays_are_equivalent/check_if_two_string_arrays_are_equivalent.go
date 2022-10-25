package check_if_two_string_arrays_are_equivalent

func ArrayStringsAreEqual(word1 []string, word2 []string) bool {
	var (
		word1Idx int
		word2Idx int
		char1Idx int
		char2Idx int
	)

	for word1Idx < len(word1) && word2Idx < len(word2) {
		if word1[word1Idx][char1Idx] != word2[word2Idx][char2Idx] {
			return false
		}

		char1Idx++
		char2Idx++

		if char1Idx == len(word1[word1Idx]) {
			word1Idx++
			char1Idx = 0
		}

		if char2Idx == len(word2[word2Idx]) {
			word2Idx++
			char2Idx = 0
		}
	}

	return word1Idx == len(word1) && word2Idx == len(word2)
}
