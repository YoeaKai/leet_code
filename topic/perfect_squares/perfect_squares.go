package perfect_squares

import "math"

func NumSquares(n int) int {
	if isSquare(n) {
		return 1
	}

	for (n & 3) == 0 {
		n >>= 2
	}

	if (n & 7) == 7 {
		return 4
	}

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
