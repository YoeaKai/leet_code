package insert_interval

func Insert(intervals [][]int, newInterval []int) [][]int {
	idx := 0

	for idx < len(intervals) && intervals[idx][1] < newInterval[0] {
		idx++
	}

	insertIdx := idx

	for ; idx < len(intervals) && intervals[idx][0] <= newInterval[1]; idx++ {
		if intervals[idx][0] < newInterval[0] {
			newInterval[0] = intervals[idx][0]
		}

		if intervals[idx][1] > newInterval[1] {
			newInterval[1] = intervals[idx][1]
		}
	}

	rest := make([][]int, len(intervals)-idx)
	copy(rest, intervals[idx:])

	return append(append(intervals[:insertIdx], newInterval), rest...)
}
