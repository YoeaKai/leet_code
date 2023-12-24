package minimum_changes_to_make_alternating_binary_string

func MinOperations(s string) int {
	zero, one := 0, 0

	for i := 0; i < len(s); i++ {
		if i%2 == 0 && s[i] == '0' || i%2 == 1 && s[i] == '1' {
			one++
		} else {
			zero++
		}
	}

	if zero < one {
		return zero
	}

	return one
}
