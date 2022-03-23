package remove_duplicates_from_sorted_array_II

func RemoveDuplicates(nums []int) int {
	idx := 1

	for i := 1; i < len(nums); i++ {
		if idx < 2 || nums[i] > nums[idx-2] {
			nums[idx] = nums[i]
			idx++
		}
	}

	return idx
}
