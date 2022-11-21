package erect_the_fence

import "sort"

func OuterTrees(trees [][]int) [][]int {
	sort.Slice(trees, func(a, b int) bool {
		if trees[a][0] == trees[b][0] {
			return trees[a][1] < trees[b][1]
		} else {
			return trees[a][0] < trees[b][0]
		}
	})

	var upper [][2]int
	var lower [][2]int

	for i := 0; i < len(trees); i++ {
		// Turn tree into [2]int data type used to be the key of the set after.
		tree := [2]int{trees[i][0], trees[i][1]}

		for len(upper) >= 2 && compare(&upper[len(upper)-2], &upper[len(upper)-1], &tree) < 0 {
			upper = upper[:len(upper)-1]
		}

		for len(lower) >= 2 && compare(&lower[len(lower)-2], &lower[len(lower)-1], &tree) > 0 {
			lower = lower[:len(lower)-1]
		}

		upper = append(upper, tree)
		lower = append(lower, tree)
	}

	set := make(map[[2]int]struct{})

	for _, tree := range upper {
		set[tree] = struct{}{}
	}

	for _, tree := range lower {
		set[tree] = struct{}{}
	}

	var res [][]int

	for tree := range set {
		res = append(res, []int{tree[0], tree[1]})
	}

	return res
}

// compare returns the value according to basic algebra below calculation for comparing
// positive and negative numbers.
// That can help us to find the angle used to judge whether it is a convex hull or not.
//
// (y3 - y2) / (x3 - x2) = slope2
// (y2 - y1) / (x2 - x1) = slope1
// slope2 - slope1 > 0
// =>
// (y3 - y2)(x2 - x1) - (x3 - x2)(y2 - y1)
// ---------------------------------------  > 0
//          (x3 - x2) * (x2 - x1)
// =>
// Because of sorting, we could get rid of the denominator.
// (y3 - y2)(x2 - x1) - (x3 - x2)(y2 - y1) > 0
// The directions are reversed between the upper and lower trees in the circle.
// It affects the greater than less than sign of the inequality.
//
// Refer: https://leetcode.com/problems/erect-the-fence/discuss/1442266/A-Detailed-Explanation-with-Diagrams-(Graham-Scan)
func compare(p1, p2, p3 *[2]int) int {
	return ((*p3)[1]-(*p2)[1])*((*p2)[0]-(*p1)[0]) - ((*p2)[1]-(*p1)[1])*((*p3)[0]-(*p2)[0])
}
