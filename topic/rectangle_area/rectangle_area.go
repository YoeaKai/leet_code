package rectangle_area

func ComputeArea(ax1 int, ay1 int, ax2 int, ay2 int, bx1 int, by1 int, bx2 int, by2 int) int {
	left := max(ax1, bx1)
	right := max(min(ax2, bx2), left)
	bottom := max(ay1, by1)
	top := max(min(ay2, by2), bottom)
	return (ax2-ax1)*(ay2-ay1) + (bx2-bx1)*(by2-by1) - (right-left)*(top-bottom)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
