package subsets

func Subsets(nums []int) [][]int {
	subsets := [][]int{{}}

	for _, num := range nums {
		for _, subset := range subsets {
			subsets = append(subsets, append(append([]int{}, subset...), num))
		}
	}

	return subsets
}
