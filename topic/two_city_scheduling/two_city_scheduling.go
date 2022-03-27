package two_city_scheduling

import "sort"

func TwoCitySchedCost(costs [][]int) (res int) {
	diff := make([]int, len(costs))

	for i, cost := range costs {
		res += cost[0]
		diff[i] = cost[1] - cost[0]
	}

	sort.Ints(diff)

	for i := 0; i < len(costs)/2; i++ {
		res += diff[i]
	}

	return res
}
