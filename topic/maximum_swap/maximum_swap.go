package maximum_swap

import "math"

type Digit struct {
	Idx int
	Val int
}

func MaximumSwap(num int) int {
	var idx int
	var head, tmp, max Digit
	nowNum := num

	for nowNum != 0 {
		digit := nowNum % 10

		if digit > tmp.Val {
			tmp.Val = digit
			tmp.Idx = idx
		} else if digit < tmp.Val {
			head.Val = digit
			head.Idx = idx
			max = tmp
		}

		nowNum /= 10
		idx++
	}

	// Decreasing numbers.
	if head.Idx == 0 {
		return num
	}

	return num + (max.Val-head.Val)*int(math.Pow(10, float64(head.Idx))) + (head.Val-max.Val)*int(math.Pow(10, float64(max.Idx)))
}
