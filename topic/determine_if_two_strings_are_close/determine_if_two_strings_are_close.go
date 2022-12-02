package determine_if_two_strings_are_close

import "sort"

func CloseStrings(word1 string, word2 string) bool {
	var exist1 int
	var exist2 int
	frequency1 := make([]int, 26)
	frequency2 := make([]int, 26)

	for _, char := range word1 {
		exist1 |= 1 << (char - 'a')
		frequency1[char-'a']++
	}

	for _, char := range word2 {
		exist2 |= 1 << (char - 'a')
		frequency2[char-'a']++
	}

	sort.Ints(frequency1)
	sort.Ints(frequency2)

	for i := 0; i < len(frequency1); i++ {
		if frequency1[i] != frequency2[i] {
			return false
		}
	}

	return exist1 == exist2
}
