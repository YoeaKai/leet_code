package concatenation_of_consecutive_binary_numbers

func ConcatenatedBinary(n int) (sum int) {
	MOD := 1000000007
	length := 0

	for i := 1; i <= n; i++ {
		if (i & (i - 1)) == 0 { // If found 100...00, then add one more number of digits.
			length++
		}

		sum = ((sum << length) | i) % MOD // Shift result and add the new number.
	}

	return sum
}
