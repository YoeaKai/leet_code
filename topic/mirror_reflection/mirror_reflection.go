package mirror_reflection

func MirrorReflection(p int, q int) int {
	for p%2 == 0 && q%2 == 0 {
		p /= 2
		q /= 2
	}

	if p%2 == 1 {
		if q%2 == 0 {
			return 0
		} else {
			return 1
		}
	}

	if p%2 == 0 && q%2 == 1 {
		return 2
	}

	return -1
}
