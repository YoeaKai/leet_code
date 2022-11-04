package reverse_vowels_of_a_string

func ReverseVowels(s string) string {
	out := []byte(s)
	l := 0
	r := len(s) - 1

	for l < r {
		for l < r && !isVowel(out[l]) {
			l++
		}
		for l < r && !isVowel(out[r]) {
			r--
		}
		if l >= r {
			break
		}
		out[l], out[r] = out[r], out[l]
		l++
		r--
	}

	return string(out)
}

func isVowel(b byte) bool {
	return b == 'a' || b == 'e' || b == 'i' || b == 'u' || b == 'o' || b == 'A' || b == 'E' || b == 'I' || b == 'O' || b == 'U'
}

func ReverseVowels2(s string) string {
	var vowels []int

	for idx, val := range s {
		if val == 'a' || val == 'e' || val == 'i' || val == 'o' || val == 'u' || val == 'A' || val == 'E' || val == 'I' || val == 'O' || val == 'U' {
			vowels = append(vowels, idx)
		}
	}

	length := len(vowels)

	for i := 0; i < length/2; i++ {
		s = s[:vowels[i]] + s[vowels[length-1-i]:vowels[length-1-i]+1] + s[vowels[i]+1:vowels[length-1-i]] + s[vowels[i]:vowels[i]+1] + s[vowels[length-1-i]+1:]
	}

	return s
}
