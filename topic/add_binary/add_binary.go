package add_binary

func AddBinary(a string, b string) string {
	if len(a) < len(b) {
		a, b = b, a
	}

	var carry byte
	res := make([]byte, len(a)+1)

	for i, j := len(a), len(b); i > 0 || j > 0; {
		var aDigit, bDigit byte

		if i > 0 {
			i -= 1
			aDigit = a[i] - '0'
		}

		if j > 0 {
			j -= 1
			bDigit = b[j] - '0'
		}

		sum := aDigit ^ bDigit ^ carry
		carry = aDigit&bDigit | carry&(aDigit^bDigit)

		res[i+1] = sum + '0'
	}

	if carry == 1 {
		res[0] = '1'
		return string(res)
	}

	return string(res[1:])
}
