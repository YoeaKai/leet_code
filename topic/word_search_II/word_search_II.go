package word_search_II

type Trie struct {
	Nodes [26]*Trie
	Word  string
}

func FindWords(board [][]byte, words []string) (res []string) {
	root := buildTrie(words)

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			dfs(board, root, i, j, &res)
		}
	}

	return res
}

func buildTrie(words []string) *Trie {
	root := &Trie{}

	for _, word := range words {
		tmp := root
		for _, c := range word {
			i := c - 'a'
			if tmp.Nodes[i] == nil {
				tmp.Nodes[i] = &Trie{}
			}
			tmp = tmp.Nodes[i]
		}
		tmp.Word = word
	}

	return root
}

func dfs(board [][]byte, root *Trie, i, j int, res *[]string) {
	if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) {
		return
	}

	tmp := board[i][j]
	if tmp == '-' || root.Nodes[tmp-'a'] == nil {
		return
	}

	root = root.Nodes[tmp-'a']
	if len(root.Word) > 0 {
		*res = append(*res, root.Word)
		root.Word = ""
	}

	board[i][j] = '-'
	dfs(board, root, i+1, j, res)
	dfs(board, root, i-1, j, res)
	dfs(board, root, i, j+1, res)
	dfs(board, root, i, j-1, res)
	board[i][j] = tmp
}
