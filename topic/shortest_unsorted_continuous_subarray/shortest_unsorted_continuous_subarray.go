package shortest_unsorted_continuous_subarray

func FindUnsortedSubarray(nums []int) int {
	n := len(nums)
	start := -1
	end := -2
	min := nums[n-1]
	max := nums[0]

	for i := 1; i < n; i++ {
		if nums[i] > max {
			max = nums[i]
		}

		if nums[n-1-i] < min {
			min = nums[n-1-i]
		}

		// Find the last index that doesn't be sorted in ascending order.
		if nums[i] < max {
			end = i
		}

		// Find the first index that doesn't be sorted in ascending order.
		if nums[n-1-i] > min {
			start = n - 1 - i
		}
	}

	return end - start + 1
}
