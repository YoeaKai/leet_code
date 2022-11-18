package ugly_number

func IsUgly(n int) bool {
	for n != 0 {
		if n%2 == 0 {
			n /= 2
			continue
		}

		if n%3 == 0 {
			n /= 3
			continue
		}

		if n%5 == 0 {
			n /= 5
			continue
		}

		break
	}

	return n == 1
}
