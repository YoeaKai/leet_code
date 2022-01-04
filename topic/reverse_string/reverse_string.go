package reverse_string

func ReverseString(s []byte) {
	front := 0
	tail := len(s) - 1

	for front < tail {
		s[front], s[tail] = s[tail], s[front]
		front++
		tail--
	}
}
