package toeplitz_matrix

func IsToeplitzMatrix(matrix [][]int) bool {
	for i := 0; i < len(matrix)-1; i++ {
		for j := 0; j < len(matrix[i])-1; j++ {
			if matrix[i][j] != matrix[i+1][j+1] {
				return false
			}
		}
	}
	return true
}

func IsToeplitzMatrix2(matrix [][]int) bool {
	n := len(matrix)
	m := len(matrix[0])

	for i := 0; i < n; i++ {
		val := matrix[i][0]
		for j := 1; j < m && i+j < n; j++ {
			if matrix[i+j][j] != val {
				return false
			}
		}
	}

	for i := 1; i < m; i++ {
		val := matrix[0][i]
		for j := 1; j < n && i+j < m; j++ {
			if matrix[j][i+j] != val {
				return false
			}
		}
	}

	return true
}
