package combination_sum_II

import "sort"

func sum(target int, current, candidates []int, result *[][]int) {
	if target == 0 {
		*result = append(*result, current)
	} else {
		for i, val := range candidates {
			if i != 0 && candidates[i] == candidates[i-1] {
				continue
			} else {
				if val <= target {
					tmp := make([]int, 0, len(current))
					tmp = append(tmp, current...)
					tmp = append(tmp, val)
					sum(target-val, tmp, candidates[i+1:], result)
				} else {
					break
				}
			}
		}
	}
}

func CombinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	result := [][]int{}
	sum(target, []int{}, candidates, &result)
	return result
}
