package minimum_number_of_steps_to_make_two_strings_anagram

func MinSteps(s string, t string) int {
	var letter [26]int
	res := 0

	for i := 0; i < len(s); i++ {
		letter[int(s[i]-'a')]++
		letter[int(t[i]-'a')]--
	}

	for _, val := range letter {
		if val > 0 {
			res += val
		}
	}

	return res
}
