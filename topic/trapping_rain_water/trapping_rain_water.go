package trapping_rain_water

func Trap(height []int) int {
	frontMaxVal := height[0]
	frontMaxPos := 0
	frontPos := 1
	frontDef := []int{}
	backMaxVal := height[len(height)-1]
	backMaxPos := len(height) - 1
	backPos := len(height) - 2
	backDef := []int{}
	sum := 0

	for frontPos < backPos {
		if height[frontPos] >= frontMaxVal {
			if frontDef != nil {
				for _, val := range frontDef {
					sum += val
				}
				frontDef = nil
			}
			frontMaxVal = height[frontPos]
			frontMaxPos = frontPos
		} else {
			frontDef = append(frontDef, frontMaxVal-height[frontPos])
		}

		if height[backPos] >= backMaxVal {
			if backDef != nil {
				for _, val := range backDef {
					sum += val
				}
				backDef = nil
			}
			backMaxVal = height[backPos]
			backMaxPos = backPos
		} else {
			backDef = append(backDef, backMaxVal-height[backPos])
		}

		frontPos++
		backPos--
	}

	if frontPos == backPos {
		if height[frontPos] > frontMaxVal && height[backPos] > backMaxVal {
			for _, val := range frontDef {
				sum += val
			}
			for _, val := range backDef {
				sum += val
			}
			return sum
		}
	}

	if frontMaxVal > backMaxVal {
		for _, val := range backDef {
			sum += val
		}
		backDef = nil
		for backPos > frontMaxPos {
			if height[backPos] >= backMaxVal {
				if backDef != nil {
					for _, val := range backDef {
						sum += val
					}
					backDef = nil
				}
				backMaxVal = height[backPos]
				backMaxPos = backPos
			} else {
				backDef = append(backDef, backMaxVal-height[backPos])
			}
			backPos--
		}

		for _, val := range backDef {
			sum += val
		}
	} else {
		for _, val := range frontDef {
			sum += val
		}
		frontDef = nil
		for frontPos < backMaxPos {
			if height[frontPos] >= frontMaxVal {
				if frontDef != nil {
					for _, val := range frontDef {
						sum += val
					}
					frontDef = nil
				}
				frontMaxVal = height[frontPos]
				frontMaxPos = frontPos
			} else {
				frontDef = append(frontDef, frontMaxVal-height[frontPos])
			}
			frontPos++
		}

		for _, val := range frontDef {
			sum += val
		}
	}

	return sum
}

/*
func Trap(height []int) int {
	if len(height) <= 2 {
		return 0
	}

	peak := -1
	for i := 1; i < len(height)-1; i++ {
		if h := height[i]; peak == -1 || h > height[peak] {
			peak = i
		}
	}
	result := 0

	minEdge := height[0]
	if minEdge > height[len(height)-1] {
		minEdge = height[len(height)-1]
	}

	if height[peak] <= minEdge {
		for i := 1; i < len(height)-1; i++ {
			result += minEdge - height[i]
		}
	} else {
		result = Trap(height[0:peak+1]) + Trap(height[peak:])
	}
	return result
}
*/
