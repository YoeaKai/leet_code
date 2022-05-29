package contiguous_array

func FindMaxLength(nums []int) int {
	mp := make(map[int]int)
	mp[0] = -1

	var (
		sum     int
		longest int
	)

	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			sum += 1
		} else {
			sum += -1
		}

		idx, ok := mp[sum]
		if ok {
			if longest < i-idx {
				longest = i - idx
			}
		} else {
			mp[sum] = i
		}
	}

	return longest
}
