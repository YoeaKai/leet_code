package largest_rectangle_in_histogram

func calculateRectangle(length, width int, res *int) {
	rectangle := length * width
	if *res < rectangle {
		*res = rectangle
	}
}

func LargestRectangleArea(heights []int) int {
	res := heights[0]
	stack := make([]int, 0, len(heights))

	for i := 0; i < len(heights); i++ {
		if len(stack) == 0 || heights[stack[len(stack)-1]] <= heights[i] {
			stack = append(stack, i)
		} else {
			j := len(stack) - 1

			for j >= 0 && heights[i] < heights[stack[j]] {
				if j == 0 {
					calculateRectangle(heights[stack[j]], i, &res)
				} else {
					calculateRectangle(heights[stack[j]], i-(stack[j-1]+1), &res)
				}
				j--
			}

			stack = stack[:j+1]

			if j == -1 {
				calculateRectangle(heights[i], (i + 1), &res)
			}

			// If value same as last in stack, update stack, avoid mistakes at final checkout.
			if j != -1 && heights[stack[j]] == heights[i] {
				stack[j] = i
			} else {
				stack = append(stack, i)
			}
		}
	}

	if len(stack) != 0 {
		calculateRectangle(heights[stack[0]], len(heights), &res)
		for i := 1; i < len(stack); i++ {
			calculateRectangle(heights[stack[i]], (len(heights) - 1 - stack[i-1]), &res)
		}
	}

	return res
}
