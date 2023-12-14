package difference_between_ones_and_zeros_in_row_and_column

// Method 1
func OnesMinusZeros(grid [][]int) [][]int {
	rows := make([]int, len(grid))
	columns := make([]int, len(grid[0]))

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 1 {
				rows[i]++
				columns[j]++
			} else {
				rows[i]--
				columns[j]--
			}
		}
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			grid[i][j] = rows[i] + columns[j]
		}
	}

	return grid
}

// Method 2
func OnesMinusZeros2(grid [][]int) [][]int {
	rows := make([]int, len(grid))
	columns := make([]int, len(grid[0]))

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			rows[i] += grid[i][j]
			columns[j] += grid[i][j]
		}
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			// The number of zeros = (hight + length) - (rows[i]+columns[j])
			grid[i][j] = 2*(rows[i]+columns[j]) - (len(grid) + len(grid[i]))
		}
	}

	return grid
}
