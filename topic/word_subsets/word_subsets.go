package word_subsets

func WordSubsets(words1 []string, words2 []string) (result []string) {
	var (
		i   int
		w2  [26]int
		tmp [26]int
	)

	for _, word := range words2 {
		tmp = counter(word)

		for i = 0; i < len(w2); i++ {
			if w2[i] < tmp[i] {
				w2[i] = tmp[i]
			}
		}
	}

	for _, word := range words1 {
		tmp = counter(word)

		for i = 0; i < len(tmp); i++ {
			if tmp[i] < w2[i] {
				break
			}
		}

		if i == 26 {
			result = append(result, word)
		}
	}

	return result
}

func counter(s string) [26]int {
	count := [26]int{}

	for _, c := range s {
		count[c-'a']++
	}

	return count
}
