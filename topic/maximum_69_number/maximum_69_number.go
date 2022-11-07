package maximum_69_number

import "math"

// Method 1: log
func Maximum69Number(num int) int {
	digits := int(math.Log10(float64(num)))

	for digits >= 0 {
		if num/int(math.Pow(10, float64(digits)))%10 == 6 {
			return num + int(3*math.Pow(10, float64(digits)))
		}
		digits -= 1
	}

	return num
}

// Method 2: divide
func Maximum69Number2(num int) int {
	cur := num
	leftmostSix := -1

	for i := 0; cur > 0; i++ {
		if cur%10 == 6 {
			leftmostSix = i
		}
		cur /= 10
	}

	return num + int(3*math.Pow(10, float64(leftmostSix)))
}
