package unique_binary_search_trees

import "math"

func NumTrees(n int) int {
	if n <= 1 {
		return 1
	}

	sum := 0

	for i := 0; i < n/2; i++ {
		sum += NumTrees(i) * NumTrees(n-i-1)
	}

	if n%2 == 0 {
		return sum * 2
	} else {
		return sum*2 + int(math.Pow(float64(NumTrees(n/2)), 2))
	}
}
