package count_primes

func CountPrimes(n int) (count int) {
	notPrime := make([]bool, n)

	for i := 2; i < n; i++ {
		if !notPrime[i] {
			count++
			for j := 2; i*j < n; j++ {
				notPrime[i*j] = true
			}
		}
	}

	return count
}
