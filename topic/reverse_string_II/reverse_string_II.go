package reverse_string_II

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

func ReverseStr(s string, k int) string {
	str := []byte(s)
	pos := 0
	for ; pos+k <= len(str); pos += k * 2 {
		reverse(str[pos : pos+k])
	}
	if pos < len(str) {
		reverse(str[pos:])
	}
	return string(str)
}

// Method 2
func ReverseStr2(s string, k int) string {
	pos := 0
	for ; pos+k-1 < len(s); pos += 2 * k {
		for i := 0; i < k-i-1; i++ {
			s = s[:pos+i] + string(s[pos+k-i-1]) + s[pos+i+1:pos+k-i-1] + string(s[pos+i]) + s[pos+k-i:]
		}
	}

	if pos < len(s) {
		k = len(s) - pos
		for i := 0; i < len(s)-pos-i-1; i++ {
			s = s[:pos+i] + string(s[pos+k-i-1]) + s[pos+i+1:pos+k-i-1] + string(s[pos+i]) + s[pos+k-i:]
		}
	}

	return s
}
