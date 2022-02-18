package smallest_value_of_the_rearranged_number

import "sort"

func SmallestNumber(num int64) int64 {
	var res int64
	digits := []int{}

	for {
		d := num % 10
		num = num / 10
		digits = append(digits, int(d))
		if num == 0 {
			break
		}
	}

	sort.Ints(digits)

	if digits[0] == 0 {
		for i, val := range digits {
			if val != 0 {
				digits[0] = val
				digits[i] = 0
				break
			}
		}
	}

	for _, digit := range digits {
		res = res*10 + int64(digit)
	}

	return res
}
