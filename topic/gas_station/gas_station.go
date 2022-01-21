package gas_station

func CanCompleteCircuit(gas []int, cost []int) int {
	res := 0
	sum := 0
	dif := 0

	for i := 0; i < len(gas); i++ {
		sum += gas[i] - cost[i]

		dif += gas[i] - cost[i]
		if dif < 0 {
			dif = 0
			res = i + 1
		}
	}

	if sum < 0 {
		return -1
	}

	return res
}
