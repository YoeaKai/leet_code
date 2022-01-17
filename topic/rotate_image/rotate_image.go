package rotate_image

func Rotate(matrix [][]int) {
	// Swing left and right
	n := len(matrix[0])
	for i := 0; i < n; i++ {
		for j := 0; j < n/2; j++ {
			tmp := matrix[i][j]
			matrix[i][j] = matrix[i][n-j-1]
			matrix[i][n-j-1] = tmp
		}
	}
	// Diagonal swap
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			tmp := matrix[i][j]
			matrix[i][j] = matrix[j][i]
			matrix[j][i] = tmp
		}
	}
}
