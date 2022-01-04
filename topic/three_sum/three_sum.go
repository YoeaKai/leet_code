package three_sum

import (
	"sort"
)

func isContains(resultList [][]int, items []int) bool {
	flag := false
	for _, e := range resultList {
		flag = false
		for i, item := range items {
			if item != e[i] {
				flag = true
				break
			}
		}
		if !flag {
			return true
		}
	}
	return false
}

func ThreeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return [][]int{}
	}

	sort.Ints(nums)
	resultList := [][]int{}

	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		posS := i + 1
		posE := len(nums) - 1

		for posS < posE {
			sumNum := nums[i] + nums[posS] + nums[posE]
			if sumNum < 0 {
				posS += 1
			} else if sumNum > 0 {
				posE -= 1
			} else {
				if !isContains(resultList, []int{nums[i], nums[posS], nums[posE]}) {
					resultList = append(resultList, []int{nums[i], nums[posS], nums[posE]})
				}
				posS += 1
				posE -= 1
			}
		}
	}
	return resultList
}
