package search_insert_position

func SearchInsert(nums []int, target int) int {
	head := 0
	tail := len(nums) - 1
	var pos int

	if nums[0] >= target {
		return 0
	}
	if nums[len(nums)-1] < target {
		return len(nums)
	}

	for {
		pos = (head + tail) / 2
		if nums[pos] < target {
			if nums[pos+1] >= target {
				return pos + 1
			} else {
				head = pos
				pos = (pos + tail) / 2
			}
		} else if nums[pos] == target {
			return pos
		} else {
			tail = pos
			pos = (pos + head) / 2
		}
	}
}
