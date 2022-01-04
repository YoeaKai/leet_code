package three_sum_closest

import (
	"sort"
)

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func ThreeSumClosest(nums []int, target int) int {
	if len(nums) == 3 {
		return nums[0] + nums[1] + nums[2]
	}

	sort.Ints(nums)
	result := nums[0] + nums[1] + nums[2]
	sumNum := 0

	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		posS := i + 1
		posE := len(nums) - 1

		for posS < posE {
			sumNum = nums[i] + nums[posS] + nums[posE]

			if sumNum < target {
				posS += 1
			} else if sumNum > target {
				posE -= 1
			} else {
				return target
			}

			if abs(sumNum-target) < abs(result-target) {
				result = sumNum
			}
		}
	}
	return result
}
