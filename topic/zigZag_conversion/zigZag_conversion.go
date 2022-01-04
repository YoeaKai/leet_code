package zigZag_conversion

func Convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	result := ""
	front := 0
	back := 0

	for i := 0; i < numRows; i++ {
		for j := 0; (2*numRows-2)*j-i < len(s); j++ {
			front = (2*numRows-2)*j - i
			if (front > 0 && back != front) || front == 0 {
				result += string(s[front])
			}
			back = (2*numRows-2)*j + i
			if back < len(s) && back != front {
				result += string(s[back])
			}
		}
	}

	return result
}
