package nearest_exit_from_entrance_in_maze

func NearestExit(maze [][]byte, entrance []int) int {
	hEdge := len(maze) - 1
	wEdge := len(maze[0]) - 1
	var queue [][3]int

	if entrance[0] != 0 {
		walk(&maze, &queue, entrance[0]-1, entrance[1], 0)
	}

	if entrance[0] != hEdge {
		walk(&maze, &queue, entrance[0]+1, entrance[1], 0)
	}

	if entrance[1] != 0 {
		walk(&maze, &queue, entrance[0], entrance[1]-1, 0)
	}

	if entrance[1] != wEdge {
		walk(&maze, &queue, entrance[0], entrance[1]+1, 0)
	}

	maze[entrance[0]][entrance[1]] = '+'

	for len(queue) != 0 {
		length := len(queue)

		for i := 0; i < length; i++ {
			if queue[i][0] == 0 || queue[i][1] == 0 || queue[i][0] == hEdge || queue[i][1] == wEdge {
				return queue[i][2]
			}

			walk(&maze, &queue, queue[i][0]-1, queue[i][1], queue[i][2])
			walk(&maze, &queue, queue[i][0]+1, queue[i][1], queue[i][2])
			walk(&maze, &queue, queue[i][0], queue[i][1]-1, queue[i][2])
			walk(&maze, &queue, queue[i][0], queue[i][1]+1, queue[i][2])
		}

		queue = queue[length:]
	}

	return -1
}

func walk(maze *[][]byte, queue *[][3]int, h, w, steps int) {
	if (*maze)[h][w] == '.' {
		(*queue) = append((*queue), [3]int{h, w, steps + 1})
		(*maze)[h][w] = '+'
	}
}
