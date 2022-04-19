package zero_one_matrix

func UpdateMatrix(mat [][]int) [][]int {
	m, n := len(mat), len(mat[0])
	max := m + n

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if mat[i][j] > 0 {
				up, left := max, max

				if i != 0 {
					left = mat[i-1][j]
				}

				if j != 0 {
					up = mat[i][j-1]
				}

				mat[i][j] = min(up, left) + 1
			}
		}
	}

	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if mat[i][j] > 0 {
				down, right := max, max

				if i != m-1 {
					right = mat[i+1][j]
				}

				if j != n-1 {
					down = mat[i][j+1]
				}

				mat[i][j] = min(mat[i][j], min(down, right)+1)
			}
		}
	}

	return mat
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
