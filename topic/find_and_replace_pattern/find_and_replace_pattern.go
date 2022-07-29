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

func replacePattern(word string) [24]int {
	pattern, replaced := [24]int{}, make(map[rune]int)

	for i, c := range word {
		if _, ok := replaced[c]; !ok {
			replaced[c] = len(replaced)
		}
		pattern[i] = replaced[c]
	}

	return pattern
}

// func FindAndReplacePattern(words []string, pattern string) (result []string) {
// 	for _, word := range words {
// 		if replacePattern(pattern) == replacePattern(word) {
// 			result = append(result, word)
// 		}
// 	}
// 	return result
// }

// func replacePattern(word string) string {
// 	replace := make(map[rune]int)

// 	for _, c := range word {
// 		if _, ok := replace[c]; !ok {
// 			replace[c] = len(replace)
// 		}
// 	}

// 	for i := 0; i < len(word); i++ {
// 		word[i] = replace[word[i]] + 'a'
// 	}

// 	return word
// }
