package sort_colors

func SortColors(nums []int) {
	left := findNotRed(&nums, 0)
	right := findNotBlue(&nums, len(nums)-1)

	for i := left; i <= right; i++ {
		if nums[i] == 0 {
			nums[left], nums[i] = nums[i], nums[left]
			left = findNotRed(&nums, left+1)
			i--
		} else if nums[i] == 2 {
			nums[right], nums[i] = nums[i], nums[right]
			right = findNotBlue(&nums, right-1)
			if i == left {
				left = findNotRed(&nums, left)
				i = left
			}
			i--
		}
	}
}

func findNotRed(nums *[]int, idx int) int {
	res := idx

	for ; res < len(*nums); res++ {
		if (*nums)[res] != 0 {
			break
		}
	}

	return res
}

func findNotBlue(nums *[]int, idx int) int {
	res := idx

	for ; res >= 0; res-- {
		if (*nums)[res] != 2 {
			break
		}
	}

	return res
}
