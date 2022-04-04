package kth_smallest_number_in_multiplication_table

func FindKthNumber(m int, n int, k int) (res int) {
	var mid int
	L := 1
	R := m * n

	// Binary search.
	for L <= R {
		mid = (L + R) >> 1
		if count(m, n, mid) < k {
			L = mid + 1
		} else {
			R = mid - 1
			res = mid
		}
	}

	return res
}

// count() counts number of integers less than or equal to x in multiplication table.
func count(m, n, x int) (res int) {
	for i := 1; i <= m; i++ {
		res += min(x/i, n)
	}
	return res
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
