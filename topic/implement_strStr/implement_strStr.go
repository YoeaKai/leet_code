package pow

import "strings"

func StrStr(haystack string, needle string) int {
	return strings.Index(haystack, needle)
	// switch {
	// case len(needle) == 0:
	// 	return 0
	// default:
	// 	for i := 0; i < len(haystack)-len(needle)+1; i++ {
	// 		if haystack[i] == needle[0] {
	// 			j := 1
	// 			for ; j < len(needle); j++ {
	// 				if haystack[i+j] != needle[j] {
	// 					break
	// 				}
	// 			}
	// 			if j == len(needle) {
	// 				return i
	// 			}
	// 		}
	// 	}
	// }
	// return -1
}

/*
func StrStr(haystack string, needle string) int {
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
*/
