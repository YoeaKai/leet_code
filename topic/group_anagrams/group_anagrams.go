package group_anagrams

func GroupAnagrams(strs []string) [][]string {
	res := [][]string{}
	cache := make(map[[26]byte]int)

	for i := 0; i < len(strs); i++ {
		list := [26]byte{}

		for j := 0; j < len(strs[i]); j++ {
			list[strs[i][j]-'a']++
		}

		if idx, ok := cache[list]; ok {
			res[idx] = append(res[idx], strs[i])
		} else {
			res = append(res, []string{strs[i]})
			cache[list] = len(res) - 1
		}
	}

	return res
}
