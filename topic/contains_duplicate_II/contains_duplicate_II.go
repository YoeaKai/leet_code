package contains_duplicate_II

func ContainsNearbyDuplicate(nums []int, k int) bool {
	m := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		if val, ok := m[nums[i]]; ok && i-val <= k {
			return true
		}
		m[nums[i]] = i
	}

	return false
}
