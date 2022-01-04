package jump_game_II

func Jump(nums []int) int {
	if len(nums) == 1 {
		return 0
	}

	pos := 0
	times := 1

	for nums[pos]+pos < len(nums)-1 {
		next := 0
		nextPos := 0
		for i := 1; pos+i < len(nums) && i <= nums[pos]; i++ {
			if nums[pos+i]+i >= nextPos {
				next = i
				nextPos = nums[pos+i] + i
			}
		}
		pos += next
		times++
	}

	return times
}
