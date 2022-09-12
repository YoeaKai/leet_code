package bag_of_tokens

import "sort"

func BagOfTokensScore(tokens []int, power int) (score int) {
	sort.Ints(tokens)
	n := len(tokens)

	for i := 0; i < n; i++ {
		if score == 0 && power < tokens[i] {
			break
		}

		if power-tokens[i] >= 0 {
			power -= tokens[i]
			score++
		} else if power+tokens[n-1]-tokens[i] > 0 {
			power += (tokens[n-1] - tokens[i])
			n--
		}
	}

	return score
}
