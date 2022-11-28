package find_players_with_zero_or_one_losses

func FindWinners(matches [][]int) [][]int {
	count := make([]int, 100001)

	for _, match := range matches {
		if count[match[0]] == 0 {
			count[match[0]] = 1
		}
		if count[match[1]] == 0 {
			count[match[1]] = 2
		} else {
			count[match[1]]++
		}
	}

	res := make([][]int, 2)

	for i, count := range count {
		if count == 1 {
			res[0] = append(res[0], i)
		} else if count == 2 {
			res[1] = append(res[1], i)
		}
	}

	return res
}
