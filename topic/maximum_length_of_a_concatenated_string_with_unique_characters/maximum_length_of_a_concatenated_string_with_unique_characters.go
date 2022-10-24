package maximum_length_of_a_concatenated_string_with_unique_characters

func MaxLength(arr []string) (result int) {
	length := len(arr)
	existMap := make([]int, length)
	existDouble := make([]bool, length)

	for i := 0; i < length; i++ {
		for j := 0; j < len(arr[i]); j++ {
			if (existMap[i]>>(arr[i][j]-'a'))&1 == 1 {
				existDouble[i] = true
				break
			}
			existMap[i] |= 1 << (arr[i][j] - 'a')
		}
	}

	for i := 0; i < length; i++ {
		if existDouble[i] {
			continue
		}
		result = max(result, findMax(len(arr[i]), existMap[i], i+1, &length, &existMap, &existDouble, &arr))
	}

	return result
}

func findMax(sum, nowMap, index int, length *int, existMap *[]int, existDouble *[]bool, arr *[]string) int {
	if index >= *length {
		return sum
	}

	curSum := findMax(sum, nowMap, index+1, length, existMap, existDouble, arr)

	if !(*existDouble)[index] && nowMap&(*existMap)[index] == 0 {
		curSum = max(curSum, findMax(sum+len((*arr)[index]), nowMap|(*existMap)[index], index+1, length, existMap, existDouble, arr))
	}

	return curSum
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
