package permutations_II

import (
	"sort"
)

func dfs(nums, stack []int, visited map[int]bool, res *[][]int) {
	if len(stack) == len(nums) {
		tmp := make([]int, len(stack))
		copy(tmp, stack)
		*res = append(*res, tmp)

		return
	}
	for i, val := range nums {
		if visited[i] {
			continue
		}
		// Exclude duplicate combinations
		if i > 0 && val == nums[i-1] && !visited[i-1] {
			continue
		}
		visited[i] = true
		dfs(nums, append(stack, val), visited, res)
		visited[i] = false
	}
}

func PermuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	var (
		stack []int
		res   [][]int
	)
	visited := make(map[int]bool)
	dfs(nums, stack, visited, &res)
	return res
}
