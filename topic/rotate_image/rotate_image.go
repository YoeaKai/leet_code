package rotate_image

func findPos(i, j, carry int, matrix [][]int, visited [][]bool) {
	x, y := j, len(visited)-i-1
	tmp := matrix[x][y]
	matrix[x][y] = carry
	visited[x][y] = true
	if !visited[y][len(visited)-x-1] {
		findPos(x, y, tmp, matrix, visited)
	}
}

func Rotate(matrix [][]int) {
	n := len(matrix[0])
	visited := make([][]bool, n)
	for i := range visited {
		visited[i] = make([]bool, n)
	}

	for i := range matrix {
		for j := range matrix[0] {
			if visited[i][j] {
				continue
			}
			findPos(i, j, matrix[i][j], matrix, visited)
		}
	}
}

func RotateII(matrix [][]int) {
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
			matrix[i][j] = matrix[n-j-1][n-i-1]
			matrix[n-j-1][n-i-1] = tmp
		}
	}
}
