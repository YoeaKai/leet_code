package valid_number

import (
	"regexp"
)

// Method 1, DFA (deterministic finite automaton)
func IsNumber(s string) bool {
	var (
		num   bool
		point bool
		e     bool
	)

	for i := 0; i < len(s); i++ {
		if '0' <= s[i] && s[i] <= '9' {
			num = true
		} else if s[i] == 'e' || s[i] == 'E' {
			if e || !num {
				return false
			}
			num = false
			e = true
		} else if s[i] == '.' {
			if e || point {
				return false
			}
			point = true
		} else if s[i] == '-' || s[i] == '+' {
			if i != 0 && (s[i-1] != 'e' || s[i] == 'E') {
				return false
			}
		} else {
			return false
		}
	}

	return num
}

// Method 2, regular expression
func IsNumber2(s string) bool {
	// optionalSign _ 1. / 1.1 / .1 / 1 (number) _ e/E _ optional -/+ _ at least 1 decimal
	// 1. 0 or 1 -/1
	// 2. at least 1 digit decimal
	// 3. at least 1 digit + dot + any digit | dot + at least 1 digit | digit
	// 4. matches e or E
	// 5. -/+ at least 1
	// 6. at least 1 decimal
	// 7. 4, 5 & 6 appear 0 or 1 time
	// 8. end
	res, _ := regexp.MatchString("^[+-]?((([0-9]+\\.[0-9]*)|(\\.[0-9]+))|([0-9]+))([eE][-+]?[0-9]+)?$", s)
	return res

	// validNumber := "^[+-]?((([0-9]+\\.[0-9]*)|(\\.[0-9]+))|([0-9]+))([eE][-+]?[0-9]+)?$"
	// reg, err := regexp.Compile(validNumber)

	// if err != nil {
	// 	fmt.Println(err)
	// }

	// return reg.MatchString(s)
}
