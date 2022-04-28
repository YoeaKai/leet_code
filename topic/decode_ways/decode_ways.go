package decode_ways

func NumDecodings(s string) int {
	// pre := 2
	res := 1

	// for i := 1; i < len(s); i++ {
	// 	switch {
	// 	case s[i] == '0':
	// 		pre, res = res, pre
	// 	case s[i-1] >= '3' || (s[i-1] == '2' && s[i] <= '6'):
	// 		pre, res = res, pre+res
	// 	default:
	// 		pre, res = res, pre
	// 	}
	// }

	return res
}
