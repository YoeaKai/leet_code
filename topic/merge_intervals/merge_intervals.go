package merge_intervals

import "sort"

func Merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(a, b int) bool {
		return intervals[a][0] < intervals[b][0]
	})

	res := make([][]int, 0)
	res = append(res, intervals[0])

	for i := 1; i < len(intervals); i++ {
		interval := intervals[i]

		if interval[0] <= res[len(res)-1][1] {
			if interval[1] > res[len(res)-1][1] {
				res[len(res)-1][1] = interval[1]
			}
		} else {
			res = append(res, interval)
		}
	}

	return res
}
