package numbers_with_same_consecutive_differences

func NumsSameConsecDiff(n int, k int) []int {
	result := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	for i := 2; i <= n; i++ {
		var cur []int

		for _, val := range result {
			digit := val % 10

			if digit+k < 10 {
				cur = append(cur, val*10+digit+k)
			}

			if k > 0 && digit-k >= 0 {
				cur = append(cur, val*10+digit-k)
			}
		}

		result = cur
	}

	return result
}
