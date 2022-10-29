package earliest_possible_day_of_full_bloom

import "sort"

func EarliestFullBloom(plantTime []int, growTime []int) (res int) {
	var time [][]int

	for i := 0; i < len(plantTime); i++ {
		time = append(time, []int{growTime[i], plantTime[i]})
	}

	sort.Slice(time, func(i, j int) bool {
		return time[i][0] < time[j][0]
	})

	for _, val := range time {
		res = max(res, val[0]) + val[1]
	}

	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
