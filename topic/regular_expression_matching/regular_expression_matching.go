package regular_expression_matching

func IsMatch(s string, p string) bool {
	if p == "" {
		return s == ""
	}

	if len(p) > 1 && p[1] == '*' {
		return IsMatch(s, p[2:]) || (s != "" && (p[0] == '.' || s[0] == p[0]) && IsMatch(s[1:], p))
	} else {
		return s != "" && (p[0] == '.' || s[0] == p[0]) && IsMatch(s[1:], p[1:])
	}
}
