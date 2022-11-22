package perfect_squares

import "math"

const MaxUint = ^uint(0)
const MaxInt = int(MaxUint >> 1)

// Method 1: Dynamic programming.
func NumSquares(n int) int {
	// perfectSquares[i] means the least number of perfect square numbers which sum to i.
	perfectSquares := []int{0}

	for len(perfectSquares) <= n {
		m := len(perfectSquares)
		count := MaxInt

		for i := 1; i*i <= m; i++ {
			count = min(count, perfectSquares[m-i*i]+1)
		}

		perfectSquares = append(perfectSquares, count)
	}

	return perfectSquares[n]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Method 2: Based on Lagrange's Four Square theorem.
func NumSquares2(n int) int {
	if isSquare(n) {
		return 1
	}

	// The result is 4 if and only if n can be written in the form of 4^k*(8*m + 7).
	for (n & 3) == 0 {
		n >>= 2
	}

	if (n & 7) == 7 {
		return 4
	}

	// Check whether 2 is the result.
	sqrt := (int)(math.Sqrt(float64(n)))
	for i := 1; i <= sqrt; i++ {
		if isSquare(n - i*i) {
			return 2
		}
	}

	return 3
}

func isSquare(n int) bool {
	sqrt := (int)(math.Sqrt(float64(n)))
	return sqrt*sqrt == n
}
