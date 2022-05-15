package continuous_subarray_sum

func CheckSubarraySum(nums []int, k int) bool {
	var sum int
	remainder := map[int]int{0: -1}

	for idx, num := range nums {
		sum += num
		sum %= k

		if pre, ok := remainder[sum]; ok {
			// Represents a multiple of k that has appeared before.
			// At least two, so >= 2
			if idx-pre >= 2 {
				return true
			}
		} else {
			remainder[sum] = idx
		}
	}

	return false
}
