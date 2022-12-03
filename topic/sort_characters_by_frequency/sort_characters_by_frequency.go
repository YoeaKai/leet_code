package sort_characters_by_frequency

import "sort"

func FrequencySort(s string) string {
	frequency := make([]int, 123) // 'z' = 122
	for i := 0; i < len(s); i++ {
		frequency[s[i]]++
	}

	res := []byte(s)
	sort.Slice(res, func(i, j int) bool {
		if frequency[res[i]] == frequency[res[j]] {
			return res[i] > res[j]
		}
		return frequency[res[i]] > frequency[res[j]]
	})

	return string(res)
}
