package container_with_most_water

import "math"

func MaxArea(height []int) int {
	maxSum := 0
	front := 0
	back := len(height) - 1
	for front < back {
		maxSum = int(math.Max(float64(maxSum), math.Min(float64(height[front]), float64(height[back]))*float64(back-front)))
		if height[front] > height[back] {
			back -= 1
		} else {
			front += 1
		}
	}
	return maxSum
}
