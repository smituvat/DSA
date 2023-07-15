package main

func exist(board [][]byte, word string) bool {
	n := len(board)
	m := len(board[0])

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if board[i][j] == word[0] {
				if search(i, j, n, m, board, word, 0) {
					return true
				}
			}
		}
	}
	return false
}

func search(i int, j int, n int, m int, board [][]byte, word string, k int) bool {
	if k == len(word) {
		return true
	}
	if i < 0 || j < 0 || i == n || j == m || board[i][j] != word[k] {
		return false
	}

	board[i][j] = 0
	op1 := search(i+1, j, n, m, board, word, k+1)
	op2 := search(i-1, j, n, m, board, word, k+1)
	op3 := search(i, j+1, n, m, board, word, k+1)
	op4 := search(i, j-1, n, m, board, word, k+1)
	board[i][j] = word[k]

	return op1 || op2 || op3 || op4

}
