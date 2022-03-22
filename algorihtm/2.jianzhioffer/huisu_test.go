package __jianzhioffer

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
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

//组合
//给定两个整数 n 和 k，返回范围 [1, n] 中所有可能的 k 个数的组合。
func combine(n int, k int) [][]int {
	res := make([][]int, 0)
	ans := make([]int, 0, k)

	nList := make([]int, 0, n)
	for i := 1; i <= n; i++ {
		nList = append(nList, i)
	}
	var backTrack func([]int)

	backTrack = func(list []int) {
		if len(ans) == k {
			temp := make([]int, len(ans))
			copy(temp, ans)
			res = append(res, temp)
			return
		}
		for i, n := range list {
			ans = append(ans, n)
			temp := make([]int, len(list))
			copy(temp, list)
			temp = append([]int{}, temp[i+1:]...)

			backTrack(temp)
			ans = ans[:len(ans)-1]
		}
	}
	backTrack(nList)
	return res
}

func TestCombine(t *testing.T) {

	fmt.Println(combine(4, 2))
}

func combine1(n int, k int) [][]int {

	var chosen []int
	var ans [][]int
	var recur func(i int)
	recur = func(i int) {
		// 临界条件
		if (len(chosen) > k) || (len(chosen)+n-i+1) < k {
			return
		}
		if i == n+1 {
			ans = append(ans, append([]int{}, chosen...))

			//             ans = append(ans, make([]int, 0))
			//             for _, value := range chosen {
			//                 ans[len(ans) - 1] = append(ans[len(ans) - 1], value)

			//             }
			return
		}
		recur(i + 1)
		chosen = append(chosen, i)
		recur(i + 1)
		chosen = chosen[:len(chosen)-1]
	}
	recur(1)
	return ans

}

// 回溯
func combine2(n int, k int) [][]int {
	res := [][]int{}
	track := []int{}
	var backTrack func(n, k, start int)
	backTrack = func(n, k, start int) {
		// 终止条件
		if len(track) == k {
			// var temp []int
			temp := make([]int, k)
			copy(temp, track)
			res = append(res, temp)
		}
		for i := start; i <= n; i++ {
			track = append(track, i)
			backTrack(n, k, i+1)
			track = track[:len(track)-1] // 回溯
		}
	}
	backTrack(n, k, 1)
	return res
}

//k个数  和为n
func combinationSum3(k int, n int) [][]int {
	res := make([][]int, 0)
	ans := make([]int, 0, k)

	var combin func(sum int, startIndex int, k int)

	combin = func(sum int, startIndex int, k int) {
		if sum == n && len(ans) == k {
			//记录
			temp := make([]int, len(ans))
			fmt.Println(ans)
			copy(temp, ans)
			res = append(res, temp)
			return
		}
		if sum > n || len(ans) > k {
			return
		}
		for i := startIndex; i <= 9; i++ {
			sum += i
			ans = append(ans, i)
			combin(sum, i+1, k)
			sum = sum - i
			ans = ans[:len(ans)-1]
		}
	}
	combin(0, 1, k)
	return res
}

func TestCombin(t *testing.T) {
	combinationSum3(3, 7)
}

var digitString = map[byte]string{
	'2': "abc",
	'3': "def",
	'4': "ghi",
	'5': "jkl",
	'6': "mno",
	'7': "pqrs",
	'8': "tuv",
	'9': "wxyz",
}

func letterCombinations(digits string) []string {
	res := make([]string, 0)
	ans := make([]byte, 0)
	if digits == "" {
		return res
	}
	var combin func(startIndex int)

	combin = func(startIndex int) {
		if len(ans) == len(digits) {
			temp := make([]byte, len(ans))
			copy(temp, ans)
			res = append(res, string(temp))
			return
		}

		for i := startIndex; i < len(digits); i++ {
			tempStr := digitString[digits[i]]
			for j := 0; j < len(tempStr); j++ {
				ans = append(ans, tempStr[j])
				combin(i + 1)
				ans = ans[:len(ans)-1]
			}
		}
	}
	combin(0)
	return res
}

func combinationSum(candidates []int, target int) [][]int {
	res := make([][]int, 0)
	ans := make([]int, 0)
	var combin func(sum int, startIndex int)

	combin = func(sum int, startIndex int) {
		if sum == target {
			temp := make([]int, len(ans))
			copy(temp, ans)
			res = append(res, temp)
			return
		}
		if sum > target {
			return
		}

		for i := startIndex; i < len(candidates); i++ {
			sum := sum + candidates[i]
			ans = append(ans, candidates[i])
			combin(sum, i)
			sum = sum - candidates[i]
			ans = ans[:len(ans)-1]
		}
	}

	combin(0, 0)
	return res
}

func combinationSum2(candidates []int, target int) [][]int {
	res := make([][]int, 0)
	ans := make([]int, 0)
	sort.Ints(candidates)
	var combin func(sum int, startIndex int)

	combin = func(sum int, startIndex int) {
		if sum == target {
			temp := make([]int, len(ans))
			copy(temp, ans)
			res = append(res, temp)
			return
		}
		if sum > target {
			return
		}

		for i := startIndex; i < len(candidates); i++ {
			if i > startIndex && candidates[i] == candidates[i-1] {
				continue
			}
			sum := sum + candidates[i]
			ans = append(ans, candidates[i])
			combin(sum, i+1)
			sum = sum - candidates[i]
			ans = ans[:len(ans)-1]
		}
	}

	combin(0, 0)
	return res

}

//分割回文串
//输入：s = "aab"
//输出：[["a","a","b"],["aa","b"]]
func partition(s string) [][]string {
	res := make([][]string, 0)
	ans := make([]string, 0)

	var splitStr func(startIndex int, ans []string)

	splitStr = func(startIndex int, ans []string) {
		//递归到字符串最后，将本次递归的结果，添加到最终结果中
		if startIndex >= len(s) {
			t := make([]string, len(ans))
			copy(t, ans)
			res = append(res, t)
		}

		for i := startIndex; i < len(s); i++ {
			//开始切割，从startIndex到i这个段的str，判断是否为回文，是回文，添加到本次递归结果中，否则i++再进行截取
			if isCorrStr(s, startIndex, i) {
				ans = append(ans, s[startIndex:i+1])
			} else {
				continue
			}
			//递归
			splitStr(i+1, ans)
			//回溯
			ans = ans[:len(ans)-1]
		}
	}

	splitStr(0, ans)
	return res
}

//判断是否回文
func isCorrStr(str string, i, j int) bool {
	for i < j {
		if str[i] != str[j] {
			return false
		}
		i++
		j--
	}
	return true
}

//复原ip
func restoreIpAddresses(s string) []string {
	res := make([]string, 0)
	tempRes := make([]string, 0)

	var dfs func(startIndex int)

	dfs = func(startIndex int) {
		if startIndex >= len(s) {
			if len(tempRes) == 4 {
				res = append(res, strings.Join(tempRes, "."))
			}
			return
		}
		for i := startIndex; i < len(s); i++ {
			if corrNum(s[startIndex : i+1]) {
				tempRes = append(tempRes, s[startIndex:i+1])
			} else {
				break
			}
			dfs(i + 1)
			tempRes = tempRes[:len(tempRes)-1]
		}
	}

	dfs(0)
	return res
}

func corrNum(ipStr string) bool {
	if len(ipStr) > 1 && ipStr[0] == '0' {
		return false
	}
	ip, err := strconv.Atoi(ipStr)
	if err != nil || ip > 255 {
		return false
	}
	return true
}

func subsets(nums []int) [][]int {
	res := make([][]int, 0)
	ans := make([]int, 0)
	res = append(res, []int{})
	var dfs func(startIndex int)

	dfs = func(startIndex int) {

		if startIndex >= len(nums) {
			return
		}
		for i := startIndex; i < len(nums); i++ {
			ans = append(ans, nums[i])
			temp := make([]int, len(ans))
			copy(temp, ans)
			res = append(res, temp)
			dfs(i + 1)
			ans = ans[:len(ans)-1]
		}
	}

	dfs(0)
	return res
}

func TestUsubsets(t *testing.T) {
	subsets([]int{1, 2, 3})
}

func subsetsWithDup(nums []int) [][]int {
	res := make([][]int, 0)
	signalRes := make([]int, 0)
	res = append(res, []int{})
	sort.Ints(nums)
	var dfs func(startIndex int)

	dfs = func(startIndex int) {
		for i := startIndex; i < len(nums); i++ {
			if i > startIndex && nums[i] == nums[i-1] {
				continue
			}
			signalRes = append(signalRes, nums[i])
			temp := make([]int, len(signalRes))
			copy(temp, signalRes)
			res = append(res, temp)

			dfs(i + 1)
			signalRes = signalRes[:len(signalRes)-1]
		}
	}

	dfs(0)
	return res
}

func findSubsequences(nums []int) [][]int {
	res := make([][]int, 0)
	signalRes := make([]int, 0)

	var dfs func(startIndex int)
	dfs = func(startIndex int) {
		if len(signalRes) >= 2 {
			temp := make([]int, len(signalRes))
			copy(temp, signalRes)
			res = append(res, temp)
		}
		levelMap := make(map[int]int)
		for i := startIndex; i < len(nums); i++ {

			if (len(signalRes) > 0 && nums[i] < signalRes[len(signalRes)-1]) || levelMap[nums[i]] >= 1 {
				continue
			}
			signalRes = append(signalRes, nums[i])
			levelMap[nums[i]]++

			dfs(i + 1)
			signalRes = signalRes[:len(signalRes)-1]
		}
	}
	dfs(0)
	return res
}

func TestFind(t *testing.T) {
	findSubsequences([]int{4, 4, 3, 2, 1})
}

func findSubsequences1(nums []int) [][]int {
	var subRes []int
	var res [][]int
	backTring(0, nums, subRes, &res)
	return res
}
func backTring(startIndex int, nums, subRes []int, res *[][]int) {
	if len(subRes) > 1 {
		tmp := make([]int, len(subRes))
		copy(tmp, subRes)
		*res = append(*res, tmp)
	}
	history := [201]int{} //记录本层元素使用记录
	for i := startIndex; i < len(nums); i++ {
		//分两种情况判断：一，当前取的元素小于子集的最后一个元素，则继续寻找下一个适合的元素
		//                或者二，当前取的元素在本层已经出现过了，所以跳过该元素，继续寻找
		if len(subRes) > 0 && nums[i] < subRes[len(subRes)-1] || history[nums[i]+100] == 1 {
			continue
		}
		history[nums[i]+100] = 1 //表示本层该元素使用过了
		subRes = append(subRes, nums[i])
		backTring(i+1, nums, subRes, res)
		subRes = subRes[:len(subRes)-1]
	}
}

//全排列
func permute(nums []int) [][]int {
	res := make([][]int, 0)
	signalRes := make([]int, 0)
	signalResMap := make(map[int]bool)
	var dfs func(startIndex int)

	dfs = func(startIndex int) {
		if len(signalRes) == len(nums) {
			temp := make([]int, len(signalRes))
			copy(temp, signalRes)
			res = append(res, temp)
			return
		}
		for i := startIndex; i < len(nums); i++ {
			if signalResMap[i] {
				continue
			}
			signalRes = append(signalRes, nums[i])
			signalResMap[i] = true
			dfs(0)
			signalRes = signalRes[:len(signalRes)-1]
			signalResMap[i] = false
		}
	}

	dfs(0)
	return res
}

//全排列
var res [][]int

func permute1(nums []int) [][]int {
	res = [][]int{}
	backTrack(nums, len(nums), []int{})
	return res
}
func backTrack(nums []int, numsLen int, path []int) {
	if len(nums) == 0 {
		p := make([]int, len(path))
		copy(p, path)
		res = append(res, p)
	}
	for i := 0; i < numsLen; i++ {
		cur := nums[i]
		path = append(path, cur)
		nums = append(nums[:i], nums[i+1:]...) //直接使用切片
		backTrack(nums, len(nums), path)
		nums = append(nums[:i], append([]int{cur}, nums[i:]...)...) //回溯的时候切片也要复原，元素位置不能变
		path = path[:len(path)-1]

	}
}

//[1,1,2]   全排列 有重复数据
func permuteUnique(nums []int) [][]int {
	res := make([][]int, 0)
	path := make([]int, 0)

	vis := make(map[int]bool)

	var dfs func()
	dfs = func() {
		if len(path) == len(nums) {
			temp := make([]int, len(nums))
			copy(temp, path)
			res = append(res, temp)
		}

		used := [21]int{}
		for i, n := range nums {
			if vis[i] == true || used[nums[i]+10] == 1 {
				continue
			}
			used[nums[i]+10] = 1
			vis[i] = true
			path = append(path, n)
			dfs()
			path = path[:len(path)-1]
			vis[i] = false

		}
	}
	dfs()
	fmt.Println(res)
	return res
}

//[1,1,2]   全排列 有重复数据
func permuteUnique02(nums []int) [][]int {
	res := make([][]int, 0)
	path := make([]int, 0)
	sort.Ints(nums)

	vis := make([]bool, len(nums))
	var dfs func()
	dfs = func() {
		if len(path) == len(nums) {
			temp := make([]int, len(nums))
			copy(temp, path)
			res = append(res, temp)
		}

		for i, n := range nums {
			if i > 0 && !vis[i-1] {
				fmt.Println("1233")
			}
			if vis[i] == true {
				continue
			}
			if i > 0 && nums[i] == nums[i-1] && !vis[i-1] {
				continue
			}
			vis[i] = true
			path = append(path, n)
			dfs()
			path = path[:len(path)-1]
			vis[i] = false
		}
	}
	dfs()
	fmt.Println(res)
	return res
}

func TestPerMute(t *testing.T) {
	permuteUnique02([]int{1, 1, 2})

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
