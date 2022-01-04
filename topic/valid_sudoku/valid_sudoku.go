package valid_sudoku

func IsValidSudoku(board [][]byte) bool {
	numMap := map[string]bool{}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' {
				row := string(board[i][j]) + "r" + string(byte(i))
				column := string(board[i][j]) + "c" + string(byte(j))
				block := string(board[i][j]) + "b" + string(byte(i/3)) + string(byte(j/3))
				if numMap[row] || numMap[column] || numMap[block] {
					return false
				}
				numMap[row] = true
				numMap[column] = true
				numMap[block] = true
			}
		}
	}

	return true
}
