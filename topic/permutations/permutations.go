package permutations

func dfs(nums, tmp []int, result *[][]int) {
	if len(nums) == 1 {
		tmpStr := make([]int, len(tmp))
		copy(tmpStr, tmp)
		*result = append(*result, append(tmpStr, nums...))
	}
	for i, val := range nums {
		tmpStr := make([]int, len(nums[:i]))
		copy(tmpStr, nums[:i])
		dfs(append(tmpStr, nums[i+1:]...), append(tmp, val), result)
	}
}

func Permute(nums []int) [][]int {
	var result [][]int
	dfs(nums, []int{}, &result)
	return result
}
