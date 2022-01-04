package n_queens

func SolveNQueens(n int) [][]string {
	var columnMap uint16
	result := [][]string{}
	slash, backslash := make([]bool, 2*n), make([]bool, 2*n)
	dfs(n, 0, []string{}, columnMap, slash, backslash, &result)
	return result
}

func dfs(n, row int, store []string, columnMap uint16, slash, backslash []bool, result *[][]string) {
	if len(store) == n {
		*result = append(*result, store)
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

		s := make([]byte, n)
		for j := range s {
			s[j] = byte('.')
		}
		s[i] = byte('Q')

		tmp := make([]string, len(store)+1)
		copy(tmp, store)
		tmp[len(store)] = string(s)

		dfs(n, row+1, tmp, columnMap, slash, backslash, result)

		columnMap = columnMap ^ (1 << offset)
		slash[i+row] = false
		backslash[i-row+n] = false
	}
}
