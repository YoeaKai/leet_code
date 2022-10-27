package image_overlap

func LargestOverlap(A [][]int, B [][]int) int {
	LA := []int{}
	LB := []int{}
	n := len(A)
	count := make(map[int]int)

	for i := 0; i < n*n; i++ {
		if A[i/n][i%n] == 1 {
			LA = append(LA, i/n*100+i%n)
		}
	}

	for i := 0; i < n*n; i++ {
		if B[i/n][i%n] == 1 {
			LB = append(LB, i/n*100+i%n)
		}
	}

	for _, valA := range LA {
		for _, valB := range LB {
			count[valA-valB]++
		}
	}

	res := 0

	for _, val := range count {
		res = max(res, val)
	}

	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
