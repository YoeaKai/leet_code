package combination_sum

import "sort"

func sum(current []int, candidates []int, target int, result *[][]int) {
	if target == 0 {
		*result = append(*result, current)
	} else {
		for i, val := range candidates {
			if val <= target {
				tmp := make([]int, 0, len(current))
				tmp = append(tmp, current...)
				tmp = append(tmp, val)
				sum(tmp, candidates[i:], target-val, result)
			} else {
				break
			}
		}
	}
}

func CombinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	result := [][]int{}
	sum([]int{}, candidates, target, &result)
	return result
}
