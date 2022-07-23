package pascals_triangle_II

func GetRow(rowIndex int) []int {
	result := make([]int, rowIndex+1)
	result[0] = 1

	for i := 1; i < rowIndex+1; i++ {
		for j := i; j >= 1; j-- {
			result[j] += result[j-1]
		}
	}

	return result
}
