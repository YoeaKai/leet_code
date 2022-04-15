package maximum_product_of_word_lengths

func MaxProduct(words []string) (res int) {
	value := make([]int, len(words))

	for i := 0; i < len(words); i++ {
		for j := 0; j < len(words[i]); j++ {
			value[i] |= 1 << (words[i][j] - 'a')
		}
	}

	for i := 0; i < len(words); i++ {
		for j := i + 1; j < len(words); j++ {
			if value[i]&value[j] == 0 && len(words[i])*len(words[j]) > res {
				res = len(words[i]) * len(words[j])
			}
		}
	}

	return res
}
