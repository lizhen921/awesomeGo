package __jianzhioffer

//二维矩阵是否包含某个字符串
func exist(board [][]byte, word string) bool {

	index := 0
	for i := 0; i < len(board); i++ {
		b := board[i]
		for j := 0; j < len(b); j++ {
			if dfs(i, len(board), j, len(b), index, board, word) {
				return true
			}
		}
	}

	return false
}

func dfs(i int, len_i int, j int, len_j int, index int, board [][]byte, word string) bool {
	if i < 0 || i > len_i || j < 0 || j > len_j || board[i][j] != word[index] {
		return false
	}
	if len(word) == index+1 {
		return true
	}

	temp := board[i][j]
	board[i][j] = '0'

	if dfs(i, len_i, j+1, len_j, index+1, board, word) ||
		dfs(i, len_i, j-1, len_j, index+1, board, word) ||
		dfs(i+1, len_i, j, len_j, index+1, board, word) ||
		dfs(i-1, len_i, j, len_j, index+1, board, word) {
		return true
	}
	board[i][j] = temp
	return false
}
