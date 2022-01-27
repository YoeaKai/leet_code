package spiral_matrix

func SpiralOrder(matrix [][]int) []int {
	m, n := len(matrix), len(matrix[0])
	res := make([]int, m*n)
	idx, x, y := 0, 0, 0
	addSign := true
	m--

	for n > 0 {
		// move left and right
		{
			for i := 0; i < n; i++ {
				res[idx] = matrix[x][y]
				idx++

				if addSign {
					y++
				} else {
					y--
				}
			}

			if addSign {
				x++
				y--
			} else {
				x--
				y++
			}

			n--
		}

		if m == 0 {
			// There is nothing to add.
			break
		}

		// move up and down
		{
			for i := 0; i < m; i++ {
				res[idx] = matrix[x][y]
				idx++

				if addSign {
					x++
				} else {
					x--
				}
			}

			if addSign {
				x--
				y--
				addSign = false
			} else {
				x++
				y++
				addSign = true
			}

			m--
		}
	}

	return res
}
