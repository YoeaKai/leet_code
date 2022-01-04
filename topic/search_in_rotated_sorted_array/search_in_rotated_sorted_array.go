package search_in_rotated_sorted_array

func Search(nums []int, target int) int {
	front := 0
	pos := len(nums) / 2
	back := len(nums) - 1
	if target == nums[front] {
		return front
	}

	if target == nums[back] {
		return back
	}

	for back-front > 1 {
		if target == nums[front] {
			return front
		}

		if target == nums[back] {
			return back
		}

		if target > nums[pos] {
			if target > nums[back] {
				if nums[front] <= nums[back] {
					return -1
				} else {
					if target < nums[front] || nums[front+1] < nums[back] {
						return -1
					}
					if nums[pos] > nums[back] {
						front = pos
						pos = (pos + back) / 2
					} else {
						back = pos
						pos = (pos + front) / 2
					}
				}
			} else {
				if back-pos == 1 {
					return -1
				}
				front = pos
				pos = (pos + back) / 2
			}
		} else if target < nums[pos] {
			if target < nums[front] {
				if nums[front] <= nums[back] {
					return -1
				} else {
					if target > nums[back] || nums[back-1] > nums[back] {
						return -1
					}
					if nums[pos] < nums[front] {
						back = pos
						pos = (pos + front) / 2
					} else {
						front = pos
						pos = (pos + back) / 2
					}
				}
			} else {
				if pos-front == 1 {
					return -1
				}
				back = pos
				pos = (pos + front) / 2
			}
		} else {
			return pos
		}
	}
	return -1
}
