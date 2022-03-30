package search_in_rotated_sorted_array_II

func Search(nums []int, target int) bool {
	left, mid, right := 0, 0, len(nums)-1

	for left <= right {
		mid = (left + right) >> 1

		if nums[mid] == target {
			return true
		}

		if (nums[left] == nums[mid]) && (nums[right] == nums[mid]) {
			left++
			right--
		} else if nums[left] <= nums[mid] {
			if (nums[left] <= target) && (nums[mid] > target) {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			if (nums[mid] < target) && (nums[right] >= target) {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}

	return false
}
