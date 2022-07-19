package pascals_triangle

func Generate(numRows int) [][]int {
	result := make([][]int, numRows)

	for i := 0; i < numRows; i++ {
		result[i] = make([]int, i+1)
		result[i][0] = 1
		result[i][i] = 1
		for j := 1; j < i; j++ {
			result[i][j] = result[i-1][j-1] + result[i-1][j]
		}
	}

	return result
}

func Generate2(numRows int) [][]int {
	result := [][]int{{1}}

	for i := 1; i < numRows; i++ {
		result = append(result, []int{1})

		for j := 0; len(result) >= 2 && j < len(result[len(result)-2])-1; j++ {
			result[len(result)-1] = append(result[len(result)-1], result[len(result)-2][j]+result[len(result)-2][j+1])
		}

		result[len(result)-1] = append(result[len(result)-1], 1)
	}

	return result
}
