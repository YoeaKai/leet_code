package widest_vertical_area_between_two_points_containing_no_points

import "sort"

func MaxWidthOfVerticalArea(points [][]int) (maxWidth int) {
	sort.Slice(points, func(i, j int) bool { return points[i][0] < points[j][0] })

	for i := 1; i < len(points); i++ {
		dif := points[i][0] - points[i-1][0]

		if maxWidth < dif {
			maxWidth = dif
		}
	}

	return maxWidth
}
