package search_a_2D_matrix

func SearchMatrix(matrix [][]int, target int) bool {
	m := len(matrix[0])
	left := 0
	right := m*len(matrix) - 1

	for left != right {
		mid := (left + right - 1) >> 1
		if matrix[mid/m][mid%m] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return matrix[left/m][left%m] == target
}
