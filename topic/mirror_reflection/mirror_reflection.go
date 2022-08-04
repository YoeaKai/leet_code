package mirror_reflection

func MirrorReflection(p int, q int) int {
	for p%2 == 0 && q%2 == 0 {
		p /= 2
		q /= 2
	}

	if q == 0 {
		return 0
	}

	if p == 1 {
		return 1
	}

	return 2
}
