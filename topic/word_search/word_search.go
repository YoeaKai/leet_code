package word_search

func Exist(board [][]byte, word string) bool {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if exist(board, word, i, j) {
				return true
			}
		}
	}

	return false
}

func exist(board [][]byte, word string, i, j int) bool {
	if i < 0 || j < 0 || i >= len(board) || j >= len(board[0]) || board[i][j] != word[0] {
		return false
	} else if len(word) == 1 && word[0] == board[i][j] {
		return true
	}

	tmp := board[i][j]
	board[i][j] = '0'
	res := exist(board, word[1:], i-1, j) || exist(board, word[1:], i+1, j) || exist(board, word[1:], i, j-1) || exist(board, word[1:], i, j+1)
	board[i][j] = tmp

	return res
}
