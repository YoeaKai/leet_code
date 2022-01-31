package jump_game

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func CanJump(nums []int) bool {
	idx, reach := 0, 0
	for reach < len(nums) && idx <= reach {
		reach = max(idx+nums[idx], reach)
		idx++
	}
	return reach >= len(nums)-1
}
