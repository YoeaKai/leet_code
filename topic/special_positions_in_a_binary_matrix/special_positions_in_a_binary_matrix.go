package special_positions_in_a_binary_matrix

func NumSpecial(mat [][]int) int {
	var existsInRow []int

	for i1, row := range mat {
		existsPos := -1

		for i2, num := range row {
			if num == 1 {
				if i1 != 0 {
					mat[0][i2]++
				}

				if existsPos == -1 {
					existsPos = i2
				} else if existsPos != -2 {
					existsPos = -2
				}
			}
		}

		if existsPos >= 0 {
			existsInRow = append(existsInRow, existsPos)
		}
	}

	var result int

	for _, exists := range existsInRow {
		if mat[0][exists] == 1 {
			result++
		}
	}

	return result
}
