package climbing_stairs

func ClimbStairs(n int) int {
	if n < 4 {
		return n
	}

	one, two := 3, 2

	for i := 3; i < n; i++ {
		one, two = one+two, one
	}

	return one
}
