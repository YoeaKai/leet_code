package implement_strStr

import "strings"

// Method 1
func StrStr(haystack string, needle string) int {
	return strings.Index(haystack, needle)
}

// Method 2
func StrStr2(haystack string, needle string) int {
	switch {
	case len(needle) == 0:
		return 0
	default:
		for i := 0; i < len(haystack)-len(needle)+1; i++ {
			if haystack[i] == needle[0] {
				j := 1
				for ; j < len(needle); j++ {
					if haystack[i+j] != needle[j] {
						break
					}
				}
				if j == len(needle) {
					return i
				}
			}
		}
	}
	return -1
}

// Method 3
func StrStr3(haystack string, needle string) int {
	if needle == "" {
		return 0
	}

	for i := 0; i < len(haystack)-len(needle)+1; i++ {
		if haystack[i] == needle[0] {
			j := 1
			for ; j < len(needle); j++ {
				if haystack[i+j] != needle[j] {
					break
				}
			}
			if j == len(needle) {
				return i
			}
		}
	}

	return -1
}
