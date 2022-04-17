package minimum_swaps_to_make_sequences_increasing

func MinSwap(nums1 []int, nums2 []int) int {
	swapDp := 1 // swapDp means for the ith element in A and B, the minimum swaps if we swap A[i] and B[i]
	fixDp := 0  // fixDp means for the ith element in A and B, the minimum swaps if we DONOT swap A[i] and B[i]

	for i := 1; i < len(nums1); i++ {
		if nums1[i] <= nums2[i-1] || nums2[i] <= nums1[i-1] {
			swapDp++
			// fixDp = fixDp
		} else if nums1[i] <= nums1[i-1] || nums2[i] <= nums2[i-1] {
			tmp := swapDp
			swapDp = fixDp + 1
			fixDp = tmp
		} else {
			min := min(swapDp, fixDp)
			swapDp = min + 1
			fixDp = min
		}
	}

	return min(swapDp, fixDp)
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
