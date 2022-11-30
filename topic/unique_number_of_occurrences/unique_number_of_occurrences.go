package unique_number_of_occurrences

func UniqueOccurrences(arr []int) bool {
	count := map[int]int{}
	for _, val := range arr {
		count[val] += 1
	}

	exist := map[int]struct{}{}
	for _, val := range count {
		if _, ok := exist[val]; ok {
			return false
		}
		exist[val] = struct{}{}
	}

	return true
}
