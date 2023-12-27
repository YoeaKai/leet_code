package minimum_time_to_make_rope_colorful

func MinCost(colors string, neededTime []int) int {
	res := 0
	maxCost := neededTime[0]

	for i := 1; i < len(colors); i++ {
		if colors[i] != colors[i-1] {
			maxCost = 0
		}

		res += min(maxCost, neededTime[i])

		if maxCost < neededTime[i] {
			maxCost = neededTime[i]
		}
	}

	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
