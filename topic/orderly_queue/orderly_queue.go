import orderly_queue

func OrderlyQueue(s string, k int) string {
    str := []byte(s)

    if k > 1{
		sort.Slice(str, func(i, j int) bool {
			return str[i] < str[j]
		})
		return string(str)
    }

	min := []byte(s)

	for i := 0; i < len(s)-1; i++ {
		str = append(str[1:], str[0])
		if bytes.Compare(str, min) < 0 {
			copy(min, str)
		}
	}

	return string(min)
}