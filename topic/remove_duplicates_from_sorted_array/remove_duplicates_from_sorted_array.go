package remove_duplicates_from_sorted_array

func RemoveDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	pos := 0

	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[pos] {
			pos++
			nums[pos] = nums[i]
		}
	}

	return pos + 1
}
