// Accepting the above question, can you output the path?

package interview_SNLG_2022_10

import "errors"

func Q2(nums []int) (res []int, err error) {
	index := len(nums) - 1

	for currPos := len(nums) - 2; currPos >= 0; currPos-- {
		if nums[currPos]+currPos >= index {
			tmp := []int{index - currPos}
			tmp = append(tmp, res...)
			copy(res, tmp)
			index = currPos
		}
	}

	if index != 0 {
		return nil, errors.New("can't find a path to the end")
	}

	return res, nil
}
