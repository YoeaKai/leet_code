package most_stones_removed_with_same_row_or_column

func RemoveStones(stones [][]int) int {
	var islands int
	m := make(map[int]int)

	for i := 0; i < len(stones); i++ {
		union(stones[i][0], ^stones[i][1], &islands, &m)
	}

	return len(stones) - islands
}

func find(x int, islands *int, m *map[int]int) int {
	if _, ok := (*m)[x]; !ok {
		(*m)[x] = x
		*islands++
	}

	if x != (*m)[x] {
		(*m)[x] = find((*m)[x], islands, m)
	}

	return (*m)[x]
}

func union(x, y int, islands *int, m *map[int]int) {
	x = find(x, islands, m)
	y = find(y, islands, m)

	if x != y {
		(*m)[x] = y
		*islands--
	}
}
