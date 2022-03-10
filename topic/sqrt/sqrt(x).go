package sqrt

// Method 1, binary search
func MySqrt(x int) int {
	if x < 2 {
		return x
	}

	res := 1
	back := x

	for res < back {
		mid := (back + res) / 2
		square := mid * mid
		if square <= x && (mid+1)*(mid+1) > x {
			return mid
		} else if square > x {
			back = mid - 1
		} else {
			res = mid + 1
		}
	}

	return res
}

// Method 2, Netwons method
func MySqrt2(x int) int {
	if x == 0 {
		return 0
	}

	res := 0.5 * float64(x)

	for i := 0; i < 20; i++ {
		res = 0.5 * (res + float64(x)/res)
	}

	return int(res)
}
