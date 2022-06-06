package prime_arrangements

func NumPrimeArrangements(n int) int {
	var (
		limit         int = 1e9 + 7
		prime_numbers     = []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97}
		prime_count   int
		pos           int
	)

	for pos <= 24 && prime_numbers[pos] <= n {
		pos++
		prime_count++
	}

	res := 1

	// count the permutation of prime
	for i := 1; i <= prime_count; i++ {
		res = res * i % limit
	}

	// count the permutation of non-prime
	for i := 1; i <= (n - prime_count); i++ {
		res = res * i % limit
	}

	return res
}
