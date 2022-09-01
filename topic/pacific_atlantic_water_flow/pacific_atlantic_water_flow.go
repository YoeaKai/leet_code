package pacific_atlantic_water_flow

func PacificAtlantic(heights [][]int) (result [][]int) {
	pacific := make(map[[2]int]bool)

	for i := 0; i < len(heights[0]); i++ {
		dfs(heights, pacific, len(heights)-1, i, -1)
	}

	for i := 0; i < len(heights); i++ {
		dfs(heights, pacific, i, len(heights[0])-1, -1)
	}

	atlantic := make(map[[2]int]bool)

	for i := 0; i < len(heights[0]); i++ {
		dfs(heights, atlantic, 0, i, -1)
	}

	for i := 0; i < len(heights); i++ {
		dfs(heights, atlantic, i, 0, -1)
	}

	for key := range pacific {
		if _, ok := atlantic[key]; ok {
			result = append(result, []int{key[0], key[1]})
		}
	}

	return result
}

func dfs(heights [][]int, ocean map[[2]int]bool, row, column, prev int) {
	if row < 0 || row >= len(heights) || column < 0 || column >= len(heights[0]) {
		return
	}

	if _, ok := ocean[[2]int{row, column}]; ok {
		return
	}

	curr := heights[row][column]

	if curr < prev {
		return
	}

	ocean[[2]int{row, column}] = true

	dfs(heights, ocean, row-1, column, curr)
	dfs(heights, ocean, row+1, column, curr)
	dfs(heights, ocean, row, column-1, curr)
	dfs(heights, ocean, row, column+1, curr)
}
