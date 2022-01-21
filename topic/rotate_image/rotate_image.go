package rotate_image

// Method 1
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

// Method 2
func Rotate2(matrix [][]int) {
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

// Method 3
func Rotate3(matrix [][]int) {
	var m int
	n := len(matrix) / 2
	l := len(matrix)
	if len(matrix)%2 == 0 {
		m = n
	} else {
		m = n + 1
	}
	// Move according to different quadrants
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			x, y := i, j
			carry := matrix[i][j]
			for k := 0; k < 4; k++ {
				x, y = y, l-x-1
				tmp := matrix[x][y]
				matrix[x][y] = carry
				carry = tmp
			}
		}
	}
}
