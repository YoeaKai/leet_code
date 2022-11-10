package remove_all_adjacent_duplicates_in_string

func RemoveDuplicates(s string) string {
	str := []byte(s)
	end := 0

	for idx := 1; idx < len(str); idx++ {
		if end >= 0 && str[end] == str[idx] {
			end--
		} else {
			end++
			str[end] = str[idx]
		}
	}

	return string(str[:end+1])
}
