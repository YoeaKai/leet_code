package longest_palindromic_substring

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func LongestPalindrome(s string) string {
	str := "$#"

	for i := 0; i < len(s); i++ {
		str = str + string(s[i]) + "#"
	}

	palindromicCount := make([]int, len(str))
	center := 0
	tail := 0
	resultCenter := 0
	resultCount := 0

	for i := 1; i < len(str); i++ {
		if tail > i {
			palindromicCount[i] = min(palindromicCount[2*center-i], tail-i)
		} else {
			palindromicCount[i] = 1
		}

		for i+palindromicCount[i] < len(str) && str[i+palindromicCount[i]] == str[i-palindromicCount[i]] {
			palindromicCount[i]++
		}

		if tail < i+palindromicCount[i] {
			tail = i + palindromicCount[i]
			center = i
		}

		if resultCount < palindromicCount[i] {
			resultCount = palindromicCount[i]
			resultCenter = i
		}
	}

	return s[(resultCenter-resultCount)/2 : (resultCenter+resultCount)/2-1] // : = (resultCenter-resultCount)/2+(resultCount-1)
}
