package gray_code

import "math"

func GrayCode(n int) []int {
	res := make([]int, int(math.Pow(2, float64(n))))
	res[0], res[1] = 0, 1
	idx := 1

	for i := 1; i < n; i++ {
		idx <<= 1
		for j := 0; j < idx; j++ {
			res[idx<<1-j-1] = res[j] + idx
		}
	}

	return res
}
