package four_sum

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

func FourSum(nums []int, target int) [][]int {
	if len(nums) < 4 {
		return [][]int{}
	}

	sort.Ints(nums)
	resultList := [][]int{}
	var posS int
	var posE int

	for i := 0; i < len(nums)-3; i++ {
		for j := i + 1; j < len(nums)-2; j++ {
			posS = j + 1
			posE = len(nums) - 1
			for posS < posE {
				sumNum := nums[i] + nums[j] + nums[posS] + nums[posE]
				if sumNum < target {
					posS += 1
				} else if sumNum > target {
					posE -= 1
				} else {
					if !isContains(resultList, []int{nums[i], nums[j], nums[posS], nums[posE]}) {
						resultList = append(resultList, []int{nums[i], nums[j], nums[posS], nums[posE]})
					}
					posS += 1
					posE -= 1
				}
			}
		}
	}
	return resultList
}
