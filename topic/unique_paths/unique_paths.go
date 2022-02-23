package unique_paths

func UniquePaths(m int, n int) int {
	// Save memory
	if m < n {
		m, n = n, m
	}

	count := make([]int, n)
	count[0] = 1

	for i := 0; i < m; i++ {
		for j := 1; j < n; j++ {
			count[j] += count[j-1]
		}
	}

	return count[n-1]
}

// Method 2: DFS
