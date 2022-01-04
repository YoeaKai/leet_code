package n_queens_II

func TotalNQueens(n int) int {
	var columnMap uint16
	result := 0
	slash, backslash := make([]bool, 2*n), make([]bool, 2*n)
	dfs(n, 0, columnMap, slash, backslash, &result)
	return result
}

func dfs(n, row int, columnMap uint16, slash, backslash []bool, result *int) {
	if row == n {
		*result++
		return
	}

	for i := 0; i < n; i++ {
		offset := n - 1 - i
		if columnMap&(1<<offset) > 0 {
			continue
		}

		if slash[i+row] || backslash[i-row+n] {
			continue
		}

		columnMap |= 1 << offset
		slash[i+row] = true
		backslash[i-row+n] = true

		dfs(n, row+1, columnMap, slash, backslash, result)

		columnMap = columnMap ^ (1 << offset)
		slash[i+row] = false
		backslash[i-row+n] = false
	}
}
