package split_array_largest_sum

func SplitArray(nums []int, m int) int {
	var sum, max uint64

	for _, num := range nums {
		sum += uint64(num)
		if uint64(num) > max {
			max = uint64(num)
		}
	}

	// Binary search.
	for max < sum {
		mid := max + (sum-max)>>2
		if isValid(&nums, m, mid) {
			sum = mid
		} else {
			max = mid + 1
		}
	}

	return int(sum)
}

func isValid(nums *[]int, m int, max uint64) bool {
	var (
		sum   uint64
		count int
	)

	for _, num := range *nums {
		sum += uint64(num)
		if sum > max {
			count++
			sum = uint64(num)
			if count >= m {
				return false
			}
		}
	}

	return true
}
