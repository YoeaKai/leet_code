package binary_search

func Search(nums []int, target int) int {
	var (
		front int
		back  int = len(nums) - 1
	)

	for front < back {
		mid := (back-front)/2 + front

		if nums[mid] < target {
			front = mid + 1
		} else if nums[mid] > target {
			back = mid - 1
		} else {
			return mid
		}
	}

	if nums[front] == target {
		return front
	}

	return -1
}
