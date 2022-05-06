package nonoverlapping_Intervals

import "sort"

func EraseOverlapIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})

	end := intervals[0][1]
	count := 1

	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] >= end {
			end = intervals[i][1]
			count++
		}
	}

	return len(intervals) - count
}
