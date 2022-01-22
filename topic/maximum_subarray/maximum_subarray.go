package maximum_subarray

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func MaxSubArray(nums []int) int {
	res := nums[0]
	nowMax := nums[0]

	for i := 1; i < len(nums); i++ {
		nowMax = max(nowMax+nums[i], nums[i])
		res = max(res, nowMax)
	}

	return res
}
