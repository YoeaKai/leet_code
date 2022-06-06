package count_primes

// Method 1
func CountPrimes(n int) (count int) {
	odd := n / 2
	notPrime := make([]bool, odd)

	if n > 2 {
		count++ // add initial 2
	}

	for i := 1; i < odd; i++ {
		if !notPrime[i] {
			count++
			num := i*2 + 1
			for j := 1; (num*j-1)/2 < odd; j += 2 {
				notPrime[(num*j-1)/2] = true
			}
		}
	}

	return count
}

// Method 2
func CountPrimes2(n int) (count int) {
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
