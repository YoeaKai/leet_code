package course_schedule

func CanFinish(numCourses int, prerequisites [][]int) bool {
	graph := make(map[int][]int, numCourses)
	for _, p := range prerequisites {
		graph[p[1]] = append(graph[p[1]], p[0])
	}

	state := make([]int, numCourses) // 0 represent didn't visit
	for i := 0; i < numCourses; i++ {
		if !dfs(graph, state, i) {
			return false
		}
	}

	return true
}

func dfs(graph map[int][]int, state []int, cur int) bool {
	if state[cur] == 1 { // Visited and pass
		return true
	}

	if state[cur] == 2 { // Visited but didn't pass
		return false
	}

	state[cur] = 2
	for _, neighbor := range graph[cur] {
		if !dfs(graph, state, neighbor) { // Have circle
			return false
		}
	}
	state[cur] = 1

	return true
}
