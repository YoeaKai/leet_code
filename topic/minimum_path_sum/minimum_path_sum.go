package minimum_path_sum

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func MinPathSum(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])

	for j := 1; j < n; j++ {
		grid[0][j] += grid[0][j-1]
	}

	for i := 1; i < m; i++ {
		grid[i][0] += grid[i-1][0]
		for j := 1; j < n; j++ {
			grid[i][j] += min(grid[i-1][j], grid[i][j-1])
		}
	}

	return grid[m-1][n-1]
}
