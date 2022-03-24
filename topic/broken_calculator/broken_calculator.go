package broken_calculator

func BrokenCalc(startValue int, target int) int {
	res := 0

	for startValue < target {
		if target%2 == 0 {
			target /= 2
		} else {
			target++
		}
		res++
	}

	return res + startValue - target
}
