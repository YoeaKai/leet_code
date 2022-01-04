package count_and_say

import "fmt"

func CountAndSay(n int) string {
	if n == 1 {
		return "1"
	}

	say := CountAndSay(n - 1)
	result := ""
	count := 1

	for i := 1; i < len(say); i++ {
		if say[i] != say[i-1] {
			result = result + fmt.Sprint(count) + string(say[i-1])
			count = 1
			continue
		}
		count++
	}

	result = result + fmt.Sprint(count) + string(say[len(say)-1])

	return result
}
