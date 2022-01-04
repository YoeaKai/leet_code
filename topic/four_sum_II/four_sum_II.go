package four_sum_II

func FourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	result := 0
	countMap := make(map[int]int)

	for i := 0; i < len(nums1); i++ {
		for j := 0; j < len(nums2); j++ {
			countMap[nums1[i]+nums2[j]]++
		}
	}

	for i := 0; i < len(nums3); i++ {
		for j := 0; j < len(nums4); j++ {
			result += countMap[-1*(nums3[i]+nums4[j])]
		}
	}

	return result
}
