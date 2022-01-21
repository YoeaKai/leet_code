package stop_gninnipS_my_sdroW

func reverse(s []byte) {
	front := 0
	tail := len(s) - 1

	for front < tail {
		s[front], s[tail] = s[tail], s[front]
		front++
		tail--
	}
}

func SpinWords(str string) string {
	s := []byte(str)
	pos := 0

	for i := 0; i < len(str); i++ {
		if s[i] == ' ' {
			if i-pos >= 5 {
				reverse(s[pos:i])
			}
			pos = i + 1
		}
	}

	if len(str)-pos >= 5 {
		reverse(s[pos:len(s)])
	}

	return string(s)
}

/*
func SpinWords(str string) string {
	head := 0
	pos := 0

	for i := 0; i < len(str); i++ {
		if str[i] == ' ' {
			if pos >= 5 {
				for j := 0; head+j < i-j-1; j++ {
					str = str[:head+j] + string(str[i-j-1]) + str[head+j+1:i-j-1] + string(str[head+j]) + str[i-j:]
				}
			}
			head = i + 1
			pos = 0
		} else {
			pos++
		}
	}

	if pos >= 5 {
		for j := 0; head+j < len(str)-1-j; j++ {
			str = str[:head+j] + string(str[len(str)-1-j]) + str[head+j+1:len(str)-1-j] + string(str[head+j]) + str[len(str)-j:]
		}
	}

	return str
}
*/
