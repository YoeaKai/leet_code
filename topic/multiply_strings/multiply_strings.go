package multiply_strings

// Method 1
func Multiply(num1 string, num2 string) string {
	data := make([]byte, len(num1)+len(num2))
	for i := len(num1) - 1; i >= 0; i-- {
		for j := len(num2) - 1; j >= 0; j-- {
			num := (num1[i]-'0')*(num2[j]-'0') + data[i+j+1]
			data[i+j], data[i+j+1] = data[i+j]+num/10, num%10
		}
	}
	for i := 0; i < len(data); i++ {
		data[i] += '0'
	}
	for i, num := range data {
		if num != '0' {
			return string(data[i:])
		}
	}
	return "0"
}

// Method 2
func Multiply2(num1 string, num2 string) string {
	result := make([]byte, len(num1)+len(num2))

	for i := len(num1) - 1; i >= 0; i-- {
		for j := len(num2) - 1; j >= 0; j-- {
			num := (num1[i]-'0')*(num2[j]-'0') + result[i+j+1]
			result[i+j], result[i+j+1] = result[i+j]+num/10, num%10
		}
	}

	for i := 0; i < len(result); i++ {
		result[i] += '0'
	}

	for i, num := range result {
		if num != '0' {
			return string(result[i:])
		}
	}

	return "0"
}
