package subsets_II

import "sort"

func SubsetsWithDup(nums []int) [][]int {
	res := [][]int{{}}
	sort.Ints(nums)
	subsetsWithDup(nums, []int{}, &res)
	return res
}

func subsetsWithDup(nums []int, cur []int, res *[][]int) {
	for i := 0; i < len(nums); i++ {
		if i == 0 || nums[i] != nums[i-1] { // i == 0 means the first append the number.
			cur = append(cur, nums[i])
			subset := make([]int, len(cur))
			copy(subset, cur)
			*res = append(*res, subset)

			subsetsWithDup(nums[i+1:], cur, res)
			cur = cur[:len(cur)-1]
		}
	}
}
