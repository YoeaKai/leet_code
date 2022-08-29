package number_of_islands

func NumIslands(grid [][]byte) (count int) {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				count++
				sink(grid, i, j)
			}
		}
	}

	return count
}

func sink(grid [][]byte, i, j int) {
	if i < 0 || i > len(grid)-1 || j < 0 || j > len(grid[0])-1 || grid[i][j] == '0' {
		return
	}

	grid[i][j] = '0'

	sink(grid, i-1, j)
	sink(grid, i+1, j)
	sink(grid, i, j-1)
	sink(grid, i, j+1)
}
