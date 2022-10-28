// You are given an integer array nums.
// You are initially positioned at the array's first index,
// and each element in the array represents your maximum jump length at that position.
// Return true if you can reach the last index, or false otherwise.

package interview_SNLG_2022_10

func Q1(nums []int) bool {
	index := len(nums) - 1

	for currPos := len(nums) - 2; currPos >= 0; currPos-- {
		if nums[currPos]+currPos >= index {
			index = currPos
		}
	}

	return index == 0
}
