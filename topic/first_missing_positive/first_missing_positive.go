package first_missing_positive

func FirstMissingPositive(nums []int) int {
	nums = append(nums, 0)
	for i := 0; i < len(nums); i++ {
		if nums[i] < 0 || nums[i] >= len(nums) {
			nums[i] = 0
		}
	}
	for i := 0; i < len(nums); i++ {
		nums[nums[i]%len(nums)] += len(nums)
	}
	for i := 1; i < len(nums); i++ {
		if nums[i]/len(nums) == 0 {
			return i
		}
	}
	return len(nums)
}
