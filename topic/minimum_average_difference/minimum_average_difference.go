package minimum_average_difference

func MinimumAverageDifference(nums []int) int {
	var frontSum int
	var backSum int
	for _, num := range nums {
		frontSum += num
	}

	n := len(nums)
	res := n - 1
	min := abs(frontSum / n)

	for i := n - 1; i > 0; i-- {
		frontSum -= nums[i]
		backSum += nums[i]
		cur := abs(frontSum/i - backSum/(n-i))

		if cur <= min {
			min = cur
			res = i - 1
		}
	}

	return res
}

func abs(val int) int {
	if val > 0 {
		return val
	}
	return val * -1
}
