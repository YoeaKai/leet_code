package reverse_words_in_a_string_III

// Method 1
func reverse(s []byte) {
	front := 0
	tail := len(s) - 1

	for front < tail {
		s[front], s[tail] = s[tail], s[front]
		front++
		tail--
	}
}

func ReverseWords(s string) string {
	str := []byte(s)
	pos := 0

	for i := 0; i < len(s); i++ {
		if str[i] == ' ' {
			reverse(str[pos:i])
			pos = i + 1
		}
	}

	reverse(str[pos:])

	return string(str)
}

// Method 2
func ReverseWords2(s string) string {
	head := 0
	pos := 0

	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			for j := 0; head+j < i-j-1; j++ {
				s = s[:head+j] + string(s[i-j-1]) + s[head+j+1:i-j-1] + string(s[head+j]) + s[i-j:]
			}
			head = i + 1
			pos = 0
		} else {
			pos++
		}
	}

	for j := 0; head+j < len(s)-1-j; j++ {
		s = s[:head+j] + string(s[len(s)-1-j]) + s[head+j+1:len(s)-1-j] + string(s[head+j]) + s[len(s)-j:]
	}

	return s
}
