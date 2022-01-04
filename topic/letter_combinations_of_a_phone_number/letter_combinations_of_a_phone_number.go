package letter_combinations_of_a_phone_number

func mapping(digit byte) string {
	switch digit {
	case '2':
		return "abc"
	case '3':
		return "def"
	case '4':
		return "ghi"
	case '5':
		return "jkl"
	case '6':
		return "mno"
	case '7':
		return "pqrs"
	case '8':
		return "tuv"
	case '9':
		return "wxyz"
	default:
		return ""
	}
}

func findDigits(result []string, digits string) []string {
	str := mapping(digits[0])
	resultLen := len(result)

	for i := 0; i < len(result); i++ {
		result[i] += string(str[0])
	}

	for i := 1; i < len(str); i++ {
		for j := 0; j < resultLen; j++ {
			result = append(result, result[j][:len(result[i])-1]+string(str[i]))
		}
	}

	if len(digits) != 1 {
		return findDigits(result, digits[1:])
	}

	return result
}

func LetterCombinations(digits string) []string {
	if digits == "" {
		return []string{}
	}

	result := []string{}
	str := mapping(digits[0])

	for i := 0; i < len(str); i++ {
		result = append(result, string(str[i]))
	}

	if len(digits) == 1 {
		return result
	}

	return findDigits(result, digits[1:])
}
