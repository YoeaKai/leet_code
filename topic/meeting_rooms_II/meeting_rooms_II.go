package meeting_rooms_II

import "sort"

type Interval struct {
	Start, End int
}

func MinMeetingRooms(intervals []*Interval) int {
	start := make([]int, len(intervals))
	end := make([]int, len(intervals))

	for i, val := range intervals {
		start[i] = val.Start
		end[i] = val.End
	}

	sort.Ints(start)
	sort.Ints(end)

	i, j := 0, 0
	res := 0
	sum := 0

	for i != len(start) {
		if start[i] <= end[j] {
			i++
			sum++
		} else if start[i] > end[j] {
			j++
			sum--
		}

		if sum > res {
			res = sum
		}
	}

	return res
}
