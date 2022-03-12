package unique_paths_II

func UniquePathsWithObstacles(obstacleGrid [][]int) int {
	n := len(obstacleGrid[0])
	res := make([]int, n)
	res[0] = 1

	for _, row := range obstacleGrid {
		for i := 0; i < n; i++ {
			if row[i] == 1 {
				res[i] = 0
			} else if i > 0 {
				res[i] += res[i-1]
			}
		}
	}

	return res[n-1]
}
