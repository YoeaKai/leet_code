package valid_parentheses

func isFront(c byte) bool {
	if c == '(' || c == '[' || c == '{' {
		return true
	}
	return false
}

func reverse(c byte) byte {
	switch c {
	case '(':
		return ')'
	case ')':
		return '('
	case '{':
		return '}'
	case '}':
		return '{'
	case '[':
		return ']'
	case ']':
		return '['
	default:
		return '0'
	}
}

func IsValid(s string) bool {
	stack := []byte{s[0]}

	for i := 1; i < len(s); i++ {
		if len(stack) == 0 {
			if !isFront(s[i]) {
				return false
			}
			stack = append(stack, s[i])
		} else if stack[len(stack)-1] == reverse(s[i]) {
			if isFront(s[i]) {
				return false
			}
			stack = stack[:len(stack)-1]
		} else if isFront(s[i]) {
			stack = append(stack, s[i])
		} else {
			return false
		}
	}

	return len(stack) == 0
}
