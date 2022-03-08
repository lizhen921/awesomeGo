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

//实现 pow(x, n) ，即计算 x 的 n 次幂函数（即，xn）。不得使用库函数，同时不需要考虑大数问题。
func pow(m float64, n int) float64 {
	if m == 0 {
		return 0
	}
	res := calPow(m, n)
	if n < 0 {
		res = 1 / res
	}
	return res
}

func calPow(m float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return m
	}
	res := calPow(m, n>>1)
	res *= res
	if n&1 != n {
		res *= m
	}
	return res
}
