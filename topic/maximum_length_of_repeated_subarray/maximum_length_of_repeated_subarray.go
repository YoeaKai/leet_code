package maximum_length_of_repeated_subarray

func FindLength(nums1 []int, nums2 []int) (max int) {
	dp := make([]int, len(nums2))

	for i := 0; i < len(nums1); i++ {
		for j := len(nums2) - 1; j >= 0; j-- {
			if nums1[i] != nums2[j] {
				dp[j] = 0
			} else {
				dp[j] = 1

				if j != 0 {
					dp[j] += dp[j-1]
				}

				if max < dp[j] {
					max = dp[j]
				}
			}
		}
	}

	return max
}
