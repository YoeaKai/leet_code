package maximum_absolute_sum_of_any_subarray

// abs subarray sum
// = one prefix sum - the other prefix sum
// <= maximum prefix sum - minimum prefix sum

func MaxAbsoluteSum(nums []int) int {
	max, min, sum := 0, 0, 0

	for _, val := range nums {
		sum += val

		if sum > max {
			max = sum
		} else if sum < min {
			min = sum
		}
	}

	return max - min
}
