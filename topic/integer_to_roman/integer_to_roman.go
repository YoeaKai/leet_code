package integer_to_roman

func findChar(num, digit int) string {
	if num == 5 {
		if digit == 3 {
			return "D"
		} else if digit == 2 {
			return "L"
		} else {
			return "V"
		}
	} else {
		if digit == 4 {
			return "M"
		} else if digit == 3 {
			return "C"
		} else if digit == 2 {
			return "X"
		} else {
			return "I"
		}
	}
}

func convert(num, digit int) string {
	one := findChar(1, digit)
	five := findChar(5, digit)

	if num == 4 {
		return one + five
	} else if num == 5 {
		return five
	} else if num == 9 {
		return one + findChar(1, digit+1)
	} else {
		result := ""
		if num > 5 {
			result += five
			num -= 5
		}
		for i := 0; i < num; i++ {
			result += one
		}
		return result
	}
}

func IntegerToRoman(num int) string {
	result := ""

	for i := 1; num != 0; i++ {
		result = convert(num%10, i) + result
		num /= 10
	}

	return result
}
