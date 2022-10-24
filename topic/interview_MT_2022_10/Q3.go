//There are three kinds of islands, one grid, two grids, and three grids.
// Find the numbers of each kind of island.
// e.g.
// ".##.#"
// "#.#.."
// "#...#"
// "#.##."
// Output: [2,1,2]

package interview_MT_2022_10

func Q3(B []string) []int {
	result := make([]int, 3)

	for i := 0; i < len(B); i++ {
		for j := 0; j < len(B[i]); j++ {
			s := []rune(B[i])
			if s[j] != '.' {
				sum := dfs(&result, &B, i, j)
				result[sum-1]++
			}
		}
	}

	return result
}

func dfs(result *[]int, B *[]string, i, j int) (sum int) {
	if i < 0 || i >= len(*B) || j < 0 || j >= len((*B)[i]) {
		return 0
	}

	s := []rune((*B)[i])

	if s[j] != '.' {
		(*B)[i] = (*B)[i][:j] + "." + (*B)[i][j+1:]
		sum++
		sum += dfs(result, B, i-1, j)
		sum += dfs(result, B, i+1, j)
		sum += dfs(result, B, i, j-1)
		sum += dfs(result, B, i, j+1)
	}

	return sum
}
