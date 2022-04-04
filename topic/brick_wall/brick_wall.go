package brick_wall

func LeastBricks(wall [][]int) (res int) {
	times := make(map[int]int)

	// Count the number of crossing bricks.
	for i := 0; i < len(wall); i++ {
		tmp := -1
		for j := 0; j < len(wall[i])-1; j++ {
			tmp += wall[i][j]
			times[tmp]--
		}
	}

	for _, val := range times {
		if val < res {
			res = val
		}
	}

	return res + len(wall)
}
