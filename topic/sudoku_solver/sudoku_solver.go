package sudoku_solver

func SolveSudoku(board [][]byte) {
	solve(board, 0, 0)
}

func solve(board [][]byte, row, column int) bool {
	for i := row; i < 9; i, column = i+1, 0 {
		for j := column; j < 9; j++ {
			if board[i][j] != '.' {
				continue
			}
			for num := '1'; num <= '9'; num++ {
				if isValid(board, i, j, byte(num)) {
					board[i][j] = byte(num)
					if solve(board, i, j+1) {
						return true
					}
					board[i][j] = '.'
				}
			}
			return false
		}
	}
	return true
}

func isValid(board [][]byte, row, column int, num byte) bool {
	blockRow, blockColumn := (row/3)*3, (column/3)*3
	for i := 0; i < 9; i++ {
		if board[i][column] == num || board[row][i] == num || board[blockRow+i/3][blockColumn+i%3] == num {
			return false
		}
	}
	return true
}
