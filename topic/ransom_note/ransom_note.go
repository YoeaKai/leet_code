package ransom_note

func CanConstruct(ransomNote string, magazine string) bool {
	letterMap := make([]int, 26)

	for _, letter := range magazine {
		letterMap[letter-'a']++
	}

	for _, letter := range ransomNote {
		if letterMap[letter-'a'] == 0 {
			return false
		}
		letterMap[letter-'a']--
	}

	return true
}
