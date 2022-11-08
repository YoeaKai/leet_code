package make_the_string_great

import "math"

func MakeGood(s string) string {
	var end int
	str := []byte(s)

	for cur := 0; cur < len(str); cur++ {
		if end > 0 && math.Abs(float64(str[cur])-float64(str[end-1])) == 32 {
			// Check the front letter of the removed letter.
			end--
		} else {
			str[end] = str[cur]
			end++
		}
	}

	return string(str[:end])
}
