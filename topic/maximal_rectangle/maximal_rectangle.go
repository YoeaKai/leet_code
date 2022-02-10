package maximal_rectangle

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func MaximalRectangle(matrix [][]byte) int {
	res := 0
	m, n := len(matrix), len(matrix[0])
	hight := make([]int, n)
	// Record the leftmost index
	left := make([]int, n)
	// Record the rightmost index
	right := make([]int, n)

	for i := 0; i < n; i++ {
		right[i] = n
	}

	for i := 0; i < m; i++ {
		curLeft, curRight := 0, n
		for j := 0; j < n; j++ {
			if matrix[i][j] == '1' {
				hight[j] += 1
				// max() will find the the upper left corner of the largest rectangle.
				left[j] = max(left[j], curLeft)
			} else {
				hight[j] = 0
				left[j] = 0
				curLeft = j + 1
			}

			k := n - 1 - j

			if matrix[i][k] == '1' {
				// min() will find the the bottom right corner of largest rectangle.
				right[k] = min(right[k], curRight)
			} else {
				right[k] = n
				curRight = k
			}
		}

		for j := 0; j < n; j++ {
			// The largest rectangle in this index = (right[j] - left[j]) * hight[j]
			res = max(res, (right[j]-left[j])*hight[j])
		}
	}

	return res
}
