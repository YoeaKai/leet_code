package utf_8_validation

func ValidUtf8(data []int) bool {
	for i := 0; i < len(data); {
		if data[i] < 0x80 {
			i++
			continue
		}

		var l int

		if data[i]&0xF8 == 0xF0 {
			l = 3
		} else if data[i]&0xF0 == 0xE0 {
			l = 2
		} else if data[i]&0xE0 == 0xC0 {
			l = 1
		} else {
			return false
		}

		if i+l >= len(data) {
			return false
		}

		i++

		for tail := i + l; i < tail; i++ {
			if data[i]&0xC0 != 0x80 {
				return false
			}
		}
	}

	return true
}
