package spiral_matrix_II

func GenerateMatrix(n int) [][]int {
	res := make([][]int, n)
	for i := range res {
		res[i] = make([]int, n)
	}

	if n == 0 {
		return res
	}

	idx := 1
	rowStart := 0
	rowEnd := n - 1
	colStart := 0
	colEnd := n - 1

	for rowStart <= rowEnd && colStart <= colEnd {
		for i := colStart; i <= colEnd; i++ {
			res[rowStart][i] = idx
			idx++
		}
		rowStart++

		for i := rowStart; i <= rowEnd; i++ {
			res[i][colEnd] = idx
			idx++
		}
		colEnd--

		for i := colEnd; i >= colStart; i-- {
			if rowStart <= rowEnd {
				res[rowEnd][i] = idx
				idx++
			}
		}
		rowEnd--

		for i := rowEnd; i >= rowStart; i-- {
			if colStart <= colEnd {
				res[i][colStart] = idx
				idx++
			}
		}
		colStart++
	}

	return res
}
