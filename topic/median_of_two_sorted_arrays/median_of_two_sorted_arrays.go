package median_of_two_sorted_arrays

// Method 1
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m := len(nums1)
	n := len(nums2)

	if m > n {
		return FindMedianSortedArrays(nums2, nums1)
	}

	left := 0
	right := m

	for left <= right {
		i := (left + right) / 2
		j := (m+n+1)/2 - i

		if i > 0 && nums1[i-1] > nums2[j] {
			right = i - 1
		} else if i < m && nums2[j-1] > nums1[i] {
			left = i + 1
		} else {
			leftMax := 0
			if i == 0 {
				leftMax = nums2[j-1]
			} else if j == 0 {
				leftMax = nums1[i-1]
			} else {
				leftMax = max(nums2[j-1], nums1[i-1])
			}

			if (m+n)&1 == 1 {
				return float64(leftMax)
			}

			rightMin := 0
			if i == m {
				rightMin = nums2[j]
			} else if j == n {
				rightMin = nums1[i]
			} else {
				rightMin = min(nums2[j], nums1[i])
			}

			return float64(leftMax+rightMin) / 2
		}
	}

	return 0.0
}

// Method 2
func findMergedArr(nums1 []int, nums2 []int) []int {
	pos := 0
	head := 0
	tail := 0
	for _, val := range nums2 {
		head = 0
		tail = len(nums1) - 1
		pos = len(nums1) / 2
		for {
			if pos == len(nums1)-1 {
				if pos == 0 {
					nums1 = append([]int{val}, nums1...)
					break
				}
				if nums1[pos] > val {
					pos /= 2
					continue
				} else {
					nums1 = append(nums1, val)
					break
				}
			}
			if pos == 0 && nums1[pos] > val {
				nums1 = append([]int{val}, nums1...)
				break
			}
			if (nums1[pos] < val && nums1[pos+1] > val) || nums1[pos+1] == val || nums1[pos] == val {
				nums1 = append(nums1[:pos+1], append([]int{val}, nums1[pos+1:]...)...)
				break
			} else if nums1[pos] > val {
				tail = pos
				if pos-head == 1 {
					pos = 0
				} else {
					pos = (head + pos) / 2
				}
			} else {
				if tail-pos == 1 {
					pos = tail
				} else {
					head = pos
					pos = (pos + tail) / 2
				}
			}
		}
	}
	return nums1
}

func FindMedianSortedArrays2(nums1 []int, nums2 []int) float64 {
	var resultArr []int

	if nums1 == nil {
		resultArr = nums2
	} else if nums2 == nil {
		resultArr = nums1
	} else {
		if len(nums1) < len(nums2) {
			resultArr = findMergedArr(nums2, nums1)
		} else {
			resultArr = findMergedArr(nums1, nums2)
		}
	}

	if len(resultArr)%2 == 0 {
		return float64(resultArr[len(resultArr)/2-1]+resultArr[len(resultArr)/2]) / 2.0
	} else {
		return float64(resultArr[len(resultArr)/2])
	}
}
