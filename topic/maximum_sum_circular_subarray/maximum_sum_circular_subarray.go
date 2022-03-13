package maximum_sum_circular_subarray

func MaxSubarraySumCircular(nums []int) int {
	// If the subarray take a part of head array and a part of tail array,
	// the maximum result equals to the total sum minus the minimum subarray sum.
	var curMax, curMin, total int
	maxSum, minSum := nums[0], nums[0]

	for _, num := range nums {
		curMax = max(curMax+num, num)
		maxSum = max(curMax, maxSum)
		curMin = min(curMin+num, num)
		minSum = min(curMin, minSum)
		total += num
	}

	circularSum := total - minSum

	if maxSum > 0 && circularSum > maxSum {
		return circularSum
	}

	return maxSum
}

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
