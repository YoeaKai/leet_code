package convert_an_array_into_a_2d_array_with_conditions

func FindMatrix(nums []int) [][]int {
	count := make(map[int]int)
	maxCount := 0

	for _, num := range nums {
		count[num]++

		if count[num] > maxCount {
			maxCount = count[num]
		}
	}

	result := make([][]int, maxCount)
	for i := 0; i < maxCount; i++ {
		result[i] = make([]int, 0)
	}

	for key, value := range count {
		for i := 0; i < value; i++ {
			result[i] = append(result[i], key)
		}
	}

	return result
}
