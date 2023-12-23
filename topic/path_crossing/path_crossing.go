package path_crossing

type point struct {
	x int
	y int
}

func IsPathCrossing(path string) bool {
	visited := make(map[point]bool)
	cur := point{0, 0}
	visited[cur] = true

	for _, dir := range path {
		switch dir {
		case 'N':
			cur.y++
		case 'S':
			cur.y--
		case 'E':
			cur.x++
		case 'W':
			cur.x--
		}

		if visited[cur] {
			return true
		}

		visited[cur] = true
	}

	return false
}
