package generate_parentheses

func canAddNomber(s string) int {
	sum := 0
	for k := 0; k < len(s); k++ {
		if s[k] == '(' {
			sum++
		} else {
			sum--
		}
	}
	return sum
}

func GenerateParenthesis(n int) []string {
	result := []string{"("}
	newSlice := []string{}

	for len(result[0]) != n*2 {
		for j := 0; j < len(result); j++ {
			switch canAddNomber(result[j]) {
			case n*2 - len(result[j]):
				result[j] += ")"
			case 0:
				result[j] += "("
			default:
				newSlice = append(newSlice, result[j]+")")
				result[j] += "("
			}
		}
		result = append(result, newSlice...)
		newSlice = nil
	}

	return result
}
