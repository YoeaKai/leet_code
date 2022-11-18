package guess_number_higher_or_lower_II

const MaxUint = ^uint(0)
const MaxInt = int(MaxUint >> 1)

func GetMoneyAmount(n int) int {
	m := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		m[i] = make([]int, n+1)
	}

	return dp(&m, 1, n)
}

func dp(m *[][]int, start, end int) int {
	if start >= end {
		return 0
	}

	if (*m)[start][end] != 0 {
		return (*m)[start][end]
	}

	res := MaxInt

	for x := start; x <= end; x++ {
		cur := x + max(dp(m, start, x-1), dp(m, x+1, end))
		if cur < res {
			res = cur
		}
	}

	(*m)[start][end] = res

	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
