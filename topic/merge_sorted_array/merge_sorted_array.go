package merge_sorted_array

func Merge(nums1 []int, m int, nums2 []int, n int) {
	idx1 := m - 1
	idx2 := n - 1

	for idx := m + n - 1; idx > idx1; idx-- {
		if idx1 >= 0 && nums1[idx1] > nums2[idx2] {
			nums1[idx] = nums1[idx1]
			idx1--
		} else {
			nums1[idx] = nums2[idx2]
			idx2--
		}
	}
}
