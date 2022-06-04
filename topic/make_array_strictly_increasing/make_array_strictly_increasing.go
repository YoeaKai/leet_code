package make_array_strictly_increasing

import "sort"

const MaxUint = ^uint(0)
const MaxInt = int(MaxUint >> 1)

func MakeArrayIncreasing(arr1 []int, arr2 []int) int {
	sort.Ints(arr2)

	// First element is the number we choose for current position.
	// Second element is how many times we change the number.
	dp := make(map[int]int)
	dp[-1] = 0

	for _, v1 := range arr1 {
		tmp := make(map[int]int)

		for key, val := range dp {
			if v1 > key {
				var lastCount int

				if t, ok := tmp[v1]; ok {
					lastCount = t
				} else {
					lastCount = MaxInt
				}

				tmp[v1] = min(lastCount, val)
			}

			i := binarySearch(&arr2, key)

			if i != len(arr2) { // Remove the possibility that is not feasible.
				var lastCount int

				if t, ok := tmp[arr2[i]]; ok {
					lastCount = t
				} else {
					lastCount = MaxInt
				}

				tmp[arr2[i]] = min(lastCount, val+1)
			}
		}

		dp = tmp
	}

	if len(dp) == 0 {
		return -1
	} else {
		ret := MaxInt

		for _, v := range dp {
			ret = min(ret, v)
		}

		return ret
	}
}

func binarySearch(arr *[]int, val int) int {
	left := 0
	right := len(*arr)

	for left < right {
		mid := (left + right) / 2 // left + (right - left) / 2

		if (*arr)[mid] <= val {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return left
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
