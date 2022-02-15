package permutation_sequence

import "strconv"

func GetPermutation(n int, k int) string {
	var res int
	factorial := make([]int, n+1)
	numbers := make([]int, n+1)

	factorial[0] = 1

	for i := 1; i <= n; i++ {
		factorial[i] = factorial[i-1] * i // Create a slice of factorial.
		numbers[i-1] = i                  // Create a slice of numbers to get indices.
	}

	k--

	for i := 1; i <= n; i++ {
		// Divide the numbers into 'n-i+1' groups, each group will have 'factorial[n-i]' numbers,
		// and what we need will be the index-th number
		index := k / factorial[n-i]
		res = res*10 + numbers[index]                           // Put the number in last digit.
		numbers = append(numbers[:index], numbers[index+1:]...) // Remove the used number from numbers.
		k -= index * factorial[n-i]
	}

	return strconv.Itoa(res)
}
