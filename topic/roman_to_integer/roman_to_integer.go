package roman_to_integer

func convert(c byte) int {
	switch c {
	case 'I':
		return 1
	case 'V':
		return 5
	case 'X':
		return 10
	case 'L':
		return 50
	case 'C':
		return 100
	case 'D':
		return 500
	case 'M':
		return 1000
	default:
		return 0
	}
}

func RomanToInteger(s string) int {
	sum := 0
	num1 := 0
	num2 := 0

	for i := 0; i < len(s); i++ {
		num1 = convert(s[i])
		if i != len(s)-1 {
			num2 = convert(s[i+1])
			if num1 < num2 {
				sum += (num2 - num1)
				i++
				continue
			}
		}
		sum += num1
	}

	return sum
}
