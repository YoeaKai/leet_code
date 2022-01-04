package median_of_two_sorted_arrays

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

func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
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
