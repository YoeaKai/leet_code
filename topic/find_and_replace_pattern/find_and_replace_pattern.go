package find_and_replace_pattern

func FindAndReplacePattern(words []string, pattern string) (result []string) {
	p := replacePattern(pattern)

	for _, word := range words {
		if replacePattern(word) == p {
			result = append(result, word)
		}
	}

	return result
}

func replacePattern(word string) [50]int {
	pattern, replaced := [50]int{}, make(map[rune]int)

	for i, c := range word {
		if _, ok := replaced[c]; !ok {
			replaced[c] = len(replaced)
		}
		pattern[i] = replaced[c]
	}

	return pattern
}
