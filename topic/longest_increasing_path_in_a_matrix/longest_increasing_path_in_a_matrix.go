package longest_increasing_path_in_a_matrix

import (
	"sort"
)

// Method 1
func findVal(x, y int, matrix, maxMap *[][]int, res, max *int) {
	if (*maxMap)[x][y] == -1 {
		findMax(x, y, matrix, maxMap, res)
	}

	if (*maxMap)[x][y]+1 > *max {
		*max = (*maxMap)[x][y] + 1
	}
}

func findMax(x, y int, matrix, maxMap *[][]int, res *int) {
	(*maxMap)[x][y] = 0

	if x > 0 && (*matrix)[x-1][y] > (*matrix)[x][y] {
		findVal(x-1, y, matrix, maxMap, res, &(*maxMap)[x][y])
	}

	if x < len(*matrix)-1 && (*matrix)[x+1][y] > (*matrix)[x][y] {
		findVal(x+1, y, matrix, maxMap, res, &(*maxMap)[x][y])
	}

	if y > 0 && (*matrix)[x][y-1] > (*matrix)[x][y] {
		findVal(x, y-1, matrix, maxMap, res, &(*maxMap)[x][y])
	}

	if y < len((*matrix)[0])-1 && (*matrix)[x][y+1] > (*matrix)[x][y] {
		findVal(x, y+1, matrix, maxMap, res, &(*maxMap)[x][y])
	}
}

func LongestIncreasingPath(matrix [][]int) int {
	res := 0
	maxMap := make([][]int, len(matrix))
	for i := range maxMap {
		maxMap[i] = make([]int, len(matrix[0]))
		for j := range maxMap[i] {
			maxMap[i][j] = -1
		}
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if maxMap[i][j] == -1 {
				findMax(i, j, &matrix, &maxMap, &res)
			}
			if maxMap[i][j] > res {
				res = maxMap[i][j]
			}
		}
	}

	return res + 1
}

// Method 2
type Cell struct {
	val int
	x   int
	y   int
}

func LongestIncreasingPath2(matrix [][]int) int {
	res := 0
	maxMap := make([][]int, len(matrix))
	cellMap := make([]Cell, len(matrix)*len(matrix[0]))
	for i := 0; i < len(matrix); i++ {
		// initial maxMap
		maxMap[i] = make([]int, len(matrix[0]))

		// initial cellMap
		for j := 0; j < len(matrix[0]); j++ {
			pos := i*len(matrix[0]) + j
			cellMap[pos].val = matrix[i][j]
			cellMap[pos].x = i
			cellMap[pos].y = j
		}
	}
	sort.Slice(cellMap, func(i, j int) bool { return cellMap[i].val < cellMap[j].val })

	for _, val := range cellMap {
		if val.y > 0 && matrix[val.x][val.y-1] > matrix[val.x][val.y] && maxMap[val.x][val.y]+1 > maxMap[val.x][val.y-1] {
			maxMap[val.x][val.y-1] = maxMap[val.x][val.y] + 1
			if maxMap[val.x][val.y-1] > res {
				res = maxMap[val.x][val.y-1]
			}
		}

		if val.y < len(matrix[0])-1 && matrix[val.x][val.y+1] > matrix[val.x][val.y] && maxMap[val.x][val.y]+1 > maxMap[val.x][val.y+1] {
			maxMap[val.x][val.y+1] = maxMap[val.x][val.y] + 1
			if maxMap[val.x][val.y+1] > res {
				res = maxMap[val.x][val.y+1]
			}
		}

		if val.x > 0 && matrix[val.x-1][val.y] > matrix[val.x][val.y] && maxMap[val.x][val.y]+1 > maxMap[val.x-1][val.y] {
			maxMap[val.x-1][val.y] = maxMap[val.x][val.y] + 1
			if maxMap[val.x-1][val.y] > res {
				res = maxMap[val.x-1][val.y]
			}
		}

		if val.x < len(matrix)-1 && matrix[val.x+1][val.y] > matrix[val.x][val.y] && maxMap[val.x][val.y]+1 > maxMap[val.x+1][val.y] {
			maxMap[val.x+1][val.y] = maxMap[val.x][val.y] + 1
			if maxMap[val.x+1][val.y] > res {
				res = maxMap[val.x+1][val.y]
			}
		}
	}

	return res + 1
}
