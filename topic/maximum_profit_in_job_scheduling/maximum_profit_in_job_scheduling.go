package maximum_profit_in_job_scheduling

import "sort"

// Top-down Dynamic Programming + Binary search
func JobScheduling(startTime []int, endTime []int, profit []int) int {
	jobs := make([][3]int, len(startTime))

	for i := 0; i < len(startTime); i++ {
		jobs[i] = [3]int{startTime[i], endTime[i], profit[i]}
	}

	sort.Slice(jobs, func(i, j int) bool {
		return jobs[i][0] < jobs[j][0]
	})

	m := make(map[int]int)

	return dp(&jobs, &m, 0)
}

// dp schedules jobs from start time latest to earliest.
func dp(jobs *[][3]int, m *map[int]int, index int) int {
	if index == len(*jobs) {
		return 0
	}

	if v, ok := (*m)[index]; ok {
		return v
	}

	// Compare selecting the job of "index" and don't select the job of "index".
	maxProfit := max(dp(jobs, m, index+1), (*jobs)[index][2]+dp(jobs, m, upperBound(jobs, (*jobs)[index][1])))

	(*m)[index] = maxProfit

	return maxProfit
}

// upperBound find the index that corresponded value is the upper bound of lastEnd.
func upperBound(jobs *[][3]int, lastEnd int) int {
	front := 0
	back := len((*jobs)) - 1

	for front <= back {
		mid := (back-front)/2 + front
		if (*jobs)[mid][0] >= lastEnd {
			back = mid - 1
		} else {
			front = mid + 1
		}
	}

	return back + 1
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
