package maximize_distance_to_closest_person

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func MaxDistToClosest(seats []int) int {
	res := 0
	last := -1

	for i := 0; i < len(seats); i++ {
		if seats[i] == 1 {
			if last < 0 {
				res = i
			} else {
				res = max(res, (i-last)/2)
			}
			last = i
		}
	}

	return max(res, len(seats)-1-last)
}
