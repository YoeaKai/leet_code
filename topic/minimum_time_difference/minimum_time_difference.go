package minimum_time_difference

import (
	"sort"
	"strconv"
)

func FindMinDifference(timePoints []string) int {
	mins := timesToInts(timePoints)
	sort.Ints(mins)

	max := 1440 // Minutes in a day.
	minDiff := max

	for i := 1; i < len(mins); i++ {
		if mins[i]-mins[i-1] < minDiff {
			minDiff = mins[i] - mins[i-1]
		}
	}

	nextDay := mins[0] + max - mins[len(mins)-1]

	if nextDay < minDiff {
		return nextDay
	}

	return minDiff
}

func timesToInts(timePoints []string) []int {
	res := make([]int, len(timePoints))

	for i, t := range timePoints {
		hours, _ := strconv.Atoi(t[:2])
		mins, _ := strconv.Atoi(t[3:])
		res[i] = hours*60 + mins
	}

	return res
}
