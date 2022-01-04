package next_permutation

func NextPermutation(nums []int) {
	if len(nums) == 1 {
		return
	}

	i := len(nums) - 2

	for ; i >= 0 && nums[i] >= nums[i+1]; i-- {
	}

	if i < 0 {
		for i = 0; i < len(nums)/2; i++ {
			nums[i], nums[len(nums)-1-i] = nums[len(nums)-1-i], nums[i]
		}
		return
	}

	j := i + 1

	for ; j < len(nums) && nums[j] > nums[i]; j++ {
	}

	j--
	nums[i], nums[j] = nums[j], nums[i]
	i++

	for j = len(nums) - 1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}
