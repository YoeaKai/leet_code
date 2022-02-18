package number_of_laser_beams_in_a_bank

func NumberOfBeams(bank []string) int {
	res := 0
	prev := 0

	for row := range bank {
		one := 0

		for c := range bank[row] {
			if bank[row][c] == '1' {
				one++
			}
		}

		if one != 0 {
			res += prev * one
			prev = one
		}
	}

	return res
}
