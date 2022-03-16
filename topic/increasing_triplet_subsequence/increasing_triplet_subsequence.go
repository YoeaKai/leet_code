package increasing_triplet_subsequence

func IncreasingTriplet(nums []int) bool {
	max := int(^uint(0) >> 1) // The max value in int.
	min := nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i] < min {
			min = nums[i]
		} else if nums[i] > min {
			if nums[i] > max {
				return true
			} else if nums[i] < max {
				max = nums[i]
			}
		}
	}

	return false
}
