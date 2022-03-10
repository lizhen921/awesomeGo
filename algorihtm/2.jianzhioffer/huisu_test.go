package __jianzhioffer

import (
	"fmt"
	"testing"
)

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

func printNumbers(n int) []int {
	if n == 0 {
		return nil
	}
	max := 1
	res := make([]int, 0)
	for i := 0; i < n; i++ {
		max = max * 10
	}
	for i := 1; i < max; i++ {
		res = append(res, i)
		fmt.Println(i)
	}
	return res
}

func TestPrintNum(t *testing.T) {
	//printNumbers(3)
	printNumberStr(2)

}

func printNumberStr(n int) {
	if n == 0 {
		return
	}

	needAppend := []byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}

	res := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}
	fmt.Println(res)

	dfsNum(n, res, needAppend)

}

func dfsNum(n int, res []string, needAppend []byte) {
	//if n == 0 {
	//	return
	//}
	//for _, c := range needAppend {
	//	temp := fmt.Sprintf("%s%c", v, c)
	//	fmt.Println(temp)
	//	dfsNum(n-1, res, needAppend)
	//}
}

//全排列问题
func permute(nums []int) [][]int {
	res := make([][]int, 0)
	path := make([]int, 0, len(nums))

	var backTrack func([]int)

	backTrack = func(numbers []int) {
		if len(numbers) == 0 {
			fmt.Println("path:", path)
			temp := make([]int, len(path))
			copy(temp, path)
			res = append(res, temp)
			fmt.Println("res:", res)
			return
		}
		for i, n := range numbers {
			path = append(path, n)
			temp := make([]int, len(numbers))
			copy(temp, numbers)
			temp = append(temp[:i], temp[i+1:]...)
			backTrack(temp)
			path = path[:len(path)-1]
		}
	}

	backTrack(nums)
	fmt.Println(res)
	return res
}

func TestPerMute(t *testing.T) {
	permute([]int{1, 2, 3})

}

func permute01(nums []int) [][]int {
	//最终要返回的二维数组
	var res [][]int
	//已经用过的节点存储用的切片
	var path []int
	//将用过节点进行标记的哈希表
	visited := make(map[int]bool)
	size := len(nums)
	var backTrack func()
	backTrack = func() {
		//递归终止条件
		//也就是nums里的元素都用到了
		if len(path) == size {
			//temp暂时存放path，path的长度肯定是nums的长度
			temp := make([]int, size)
			copy(temp, path)
			res = append(res, temp)
			return
		}
		//从0开始所以不去等
		for i := 0; i < size; i++ {
			//一个排列结果（path）里的一个元素只能使用一次
			//相当于查map里有没有这个元素，有就continue跳出
			if visited[nums[i]] {
				continue
			}
			//第一次出现就给他打个标记true
			visited[nums[i]] = true
			//将这个元素放入path路径中
			path = append(path, nums[i])
			//递归
			backTrack()
			//回溯
			visited[nums[i]] = false
			//就是吐出最后一个元素
			path = path[:len(path)-1]
		}

	}
	backTrack()
	return res
}
