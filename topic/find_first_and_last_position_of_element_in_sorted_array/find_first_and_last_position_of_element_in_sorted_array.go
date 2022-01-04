package find_first_and_last_position_of_element_in_sorted_array

func SearchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}

	result := []int{-1, -1}
	i, j := 0, len(nums)

	for i < j {
		h := int(uint(i+j) >> 1)
		if target == nums[h] && (h == 0 || target != nums[h-1]) {
			result[0] = h
			break
		}
		if target > nums[h] {
			i = h + 1
		} else {
			j = h
		}
	}

	if result[0] == -1 {
		return result
	}

	if len(nums) == 1 {
		return []int{0, 0}
	}

	i, j = result[0], len(nums)

	for i < j {
		h := int(uint(i+j) >> 1)
		if target == nums[h] && (h == len(nums)-1 || target != nums[h+1]) {
			result[1] = h
			break
		}
		if target < nums[h] {
			j = h
		} else {
			i = h + 1
		}
	}

	return result
}
