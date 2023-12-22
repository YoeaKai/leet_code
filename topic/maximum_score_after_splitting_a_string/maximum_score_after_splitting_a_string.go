package maximum_score_after_splitting_a_string

func MaxScore(s string) int {
	var countOne int

	for i := 0; i < len(s); i++ {
		if s[i] == '1' {
			countOne++
		}
	}

	var maxScore int
	cur := countOne

	for i := 0; i < len(s)-1; i++ {
		if s[i] == '0' {
			cur++
		} else {
			cur--
		}

		if cur > maxScore {
			maxScore = cur
		}
	}

	return maxScore
}
