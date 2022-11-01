package where_will_the_ball_fall

func FindBall(grid [][]int) []int {
	res := make([]int, len(grid[0]))

	for i := 0; i < len(grid[0]); i++ {
		res[i] = i
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(res); j++ {
			if res[j] == -1 {
				continue
			}

			if grid[i][res[j]] == 1 {
				if res[j] == len(grid[i])-1 || grid[i][res[j]+1] == -1 {
					res[j] = -1
				} else {
					res[j]++
				}
			} else {
				if res[j] == 0 || grid[i][res[j]-1] == 1 {
					res[j] = -1
				} else {
					res[j]--
				}
			}
		}
	}

	return res
}
