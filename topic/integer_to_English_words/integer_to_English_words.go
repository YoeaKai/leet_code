package integer_to_English_words

import "strings"

func convertDigit(num int) string {
	switch num {
	case 1:
		return " One"
	case 2:
		return " Two"
	case 3:
		return " Three"
	case 4:
		return " Four"
	case 5:
		return " Five"
	case 6:
		return " Six"
	case 7:
		return " Seven"
	case 8:
		return " Eight"
	case 9:
		return " Nine"
	default:
		return ""
	}
}

func findPronunciation(num int) string {
	result := ""

	if num > 99 {
		result = convertDigit(num/100) + " Hundred"
		num = num % 100
	}

	if num/10 == 1 {
		switch num {
		case 10:
			result += " Ten"
		case 11:
			result += " Eleven"
		case 12:
			result += " Twelve"
		case 13:
			result += " Thirteen"
		case 14:
			result += " Fourteen"
		case 15:
			result += " Fifteen"
		case 16:
			result += " Sixteen"
		case 17:
			result += " Seventeen"
		case 18:
			result += " Eighteen"
		case 19:
			result += " Nineteen"
		}
	} else {
		if num/10 != 0 {
			switch num / 10 {
			case 2:
				result += " Twenty"
			case 3:
				result += " Thirty"
			case 4:
				result += " Forty"
			case 5:
				result += " Fifty"
			case 6:
				result += " Sixty"
			case 7:
				result += " Seventy"
			case 8:
				result += " Eighty"
			case 9:
				result += " Ninety"
			default:
			}
		}
		result += convertDigit(num % 10)
	}

	return result
}

func IntegerToEnglishWords(num int) string {
	if num == 0 {
		return "Zero"
	}

	result := ""

	for i := 0; num != 0; i++ {
		comma := ""

		switch i {
		case 1:
			comma = " Thousand"
		case 2:
			comma = " Million"
		case 3:
			comma = " Billion"
		}

		if num%1000 == 0 {
			num /= 1000
			continue
		}

		result = findPronunciation(num%1000) + comma + result
		num /= 1000
	}

	return strings.TrimSpace(result)
}
