package image_smoother

func ImageSmoother(img [][]int) [][]int {
	smoothed := make([][]int, len(img))
	for i := range smoothed {
		smoothed[i] = make([]int, len(img[0]))
	}

	for i := range smoothed {
		for j := range smoothed[i] {
			smoothed[i][j] = smooth(img, i, j)
		}
	}

	return smoothed
}

func smooth(img [][]int, row, col int) int {
	sum, count := 0, 0

	for i := row - 1; i <= row+1; i++ {
		for j := col - 1; j <= col+1; j++ {
			if i >= 0 && i < len(img) && j >= 0 && j < len(img[0]) {
				sum += img[i][j]
				count++
			}
		}
	}

	return sum / count
}
