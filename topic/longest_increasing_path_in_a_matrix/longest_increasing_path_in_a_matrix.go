package longest_increasing_path_in_a_matrix

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
