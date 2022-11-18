package dp

import (
	"fmt"
	"math"
	"strings"
)

func Max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func SimpleSlice(nums1 []int, nums2 []int) int {
	res := 0
	m, n := len(nums1), len(nums2)

	for i := 0; i < m+n; i++ {

	}

	return res
}

func findList(nums []int) int {

	max := 1
	tempMax := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			tempMax++
		} else if nums[i] <= nums[i-1] {
			tempMax = 1
		}
		if tempMax > max {
			max = tempMax
		}
	}
	return max
}
func findList1(nums []int) int {
	dp := make([]int, len(nums))
	max := 1
	for i := range dp {
		dp[i] = 1
	}
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = Max(dp[i], dp[j]+1)
			}
		}
		if dp[i] > max {
			fmt.Println(nums[i])
			max = dp[i]
		}
	}
	return max
}

func maxProfit3(prices []int) int {
	dp := make([][]int, len(prices))
	if len(prices) <= 1 {
		return 0
	}
	for i := 0; i < len(prices); i++ {
		dp[i] = make([]int, 1)
	}
	dp[0][0] = -prices[0]
	dp[0][1] = 0
	dp[0][2] = -prices[0]
	dp[0][3] = 0
	for i := 1; i < len(prices); i++ {
		dp[i][0] = Max(dp[i-1][0], -prices[i])
		dp[i][1] = Max(dp[i-1][1], dp[i-1][0]+prices[i])
		dp[i][2] = Max(dp[i-1][2], dp[i-1][1]-prices[i])
		dp[i][3] = Max(dp[i-1][3], dp[i-1][2]+prices[i])
	}

	return dp[len(prices)-1][3]
}

/*
持有gg-- dp[i][0] = dp[i-1][1]-price[i], dp[i-1][0]
不持有gg-- dp[i][1] = dp[i-1][0] + price[i], dp[i-1][1]
*/
func maxProfit02(prices []int) int {
	dp := make([][]int, 2)
	if len(prices) <= 1 {
		return 0
	}
	for i := 0; i < len(prices); i++ {
		dp[i] = make([]int, 1)
	}
	dp[0][0] = -prices[0]
	dp[0][1] = 0

	for i := 1; i < len(prices); i++ {
		dp[i][0] = Max(dp[i-1][1]-prices[i], dp[i-1][0])
		dp[i][1] = Max(dp[i-1][0]+prices[i], dp[i-1][1])
	}

	return dp[len(prices)-1][1]
}

func maxProfit1(prices []int) int {
	dp := make([]int, len(prices))
	if len(prices) <= 1 {
		return 0
	}
	dp[0] = 0
	min := prices[0]

	for i := 1; i < len(prices); i++ {
		dp[i] = Max(dp[i-1], prices[i]-min)
		min = Min(prices[i], min)
	}
	return dp[len(prices)-1]
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rob4(root *TreeNode) int {
	dp := robTree(root)
	return Max(dp[0], dp[1])
}

// 0--不选  1--选
func robTree(root *TreeNode) []int {
	if root == nil {
		return []int{0, 0}
	}
	left := robTree(root.Left)
	right := robTree(root.Right)
	dp3 := make([]int, 2)
	//考虑左右选，不选父节点
	dp3[0] = Max(left[0], left[1]) + Max(right[0], right[1])
	//左右不选，选父节点
	dp3[1] = root.Val + left[0] + right[0]

	return dp3
}

func rob3(nums []int) int {
	if len(nums) <= 1 {
		return nums[0]
	}

	res1 := rob2(nums[:len(nums)-1])
	res2 := rob2(nums[1:])

	return Max(res1, res2)
}

func rob2(nums []int) int {
	if len(nums) <= 1 {
		return nums[0]
	}

	a := nums[0]
	b := Max(nums[0], nums[1])

	for i := 2; i < len(nums); i++ {
		temp := Max(a+nums[i], b)
		a = b
		b = temp
	}
	return Max(a, b)
}

func rob(nums []int) int {
	dp := make([]int, len(nums))
	if len(nums) == 1 {
		return nums[0]
	}

	dp[0] = nums[0]
	dp[1] = Max(nums[0], nums[1])

	for i := 2; i < len(nums); i++ {
		dp[i] = Max(dp[i-1], dp[i-2]+nums[i])
	}
	return Max(dp[len(nums)-1], dp[len(nums)-2])
}

func wordBreak(s string, wordDict []string) bool {
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 0; i <= len(s); i++ {
		for j := 0; j < len(wordDict); j++ {
			if i >= len(wordDict[j]) && wordDict[j] == s[:i][i-len(wordDict[j]):] {
				dp[i] = dp[i] || dp[i-len(wordDict[j])]
			}
		}
	}
	return dp[len(s)]
}

/*
i 1,2,3,4,,n
n

i,j 数量
dp[i][j] = dp[i-1][j-i*i] +1
*/

func numSquares1(n int) int {
	//定义
	dp := make([]int, n+1)
	// 初始化
	dp[0] = 0
	for i := 1; i <= n; i++ {
		dp[i] = math.MaxInt32
	}
	// 遍历物品
	for i := 1; i <= n; i++ {
		// 遍历背包
		for j := i * i; j <= n; j++ {
			dp[j] = Min(dp[j], dp[j-i*i]+1)
		}
	}

	return dp[n]
}
func numSquares(n int) int {
	dp := make([]int, n+1)

	dp[0] = 0
	for i := 1; i <= n; i++ {
		dp[i] = math.MaxInt32
	}
	for i := 1; i <= n; i++ {
		for j := i * i; j <= n; j++ {
			dp[j] = Min(dp[j-i*i]+1, dp[j])
		}
	}
	if dp[n] == math.MaxInt32 {
		return -1
	}
	return dp[n]
}

func coinChange1(coins []int, amount int) int {
	dp := make([]int, amount+1)
	// 初始化dp[0]
	dp[0] = 0
	// 初始化为math.MaxInt32
	for j := 1; j <= amount; j++ {
		dp[j] = math.MaxInt32
	}

	// 遍历物品
	for i := 0; i < len(coins); i++ {
		// 遍历背包
		for j := coins[i]; j <= amount; j++ {
			if dp[j-coins[i]] != math.MaxInt32 {
				// 推导公式
				dp[j] = Min(dp[j], dp[j-coins[i]]+1)
				//fmt.Println(dp,j,i)
			}
		}
	}
	// 没找到能装满背包的, 就返回-1
	if dp[amount] == math.MaxInt32 {
		return -1
	}
	return dp[amount]
}

// 最小个数
func coinMin(coins []int, amount int) int {
	dp := make([]int, amount+1)
	//dp[0] = 0                      //背包为0是，凑齐金币肯定是0个
	for i := 0; i <= amount; i++ { //因为要求最小值，所以初始化成最大数
		dp[i] = math.MaxInt32
	}
	for i := 0; i < len(coins); i++ {
		for j := coins[i]; j <= amount; j++ {
			if dp[j-coins[i]] != math.MaxInt32 { //如果是初始值，直接跳过。因为在dp[j-coins[i]]这个位置无法凑齐背包
				dp[j] = Min(dp[j], dp[j-coins[i]]+1)
			}
		}
	}
	if dp[amount] == math.MaxInt32 {
		return -1
	}
	return int(dp[amount])
}

// 每个物品用多次，排列问题
func sum(nums []int, target int) int {
	dp := make([]int, target+1)
	dp[0] = 1
	for i := 1; i <= target; i++ {
		for j := 0; j < len(nums); j++ {
			if i < nums[j] {
				continue
			}
			dp[i] = dp[i] + dp[i-nums[j]]
		}
	}
	return dp[target]
}

func change1(coins []int, amount int) int {
	dp := make([]int, amount+1)
	dp[0] = 1
	for i := 0; i <= amount; i++ {
		for j := 0; j < len(coins); j++ {
			if i < coins[j] {
				continue
			}
			dp[i] = dp[i] + dp[i-coins[j]]
		}
	}
	return dp[amount]
}

// 零钱兑换
func change(coins []int, amount int) int {
	dp := make([]int, amount+1)
	dp[0] = 1
	for i := 0; i < len(coins); i++ {
		for j := coins[i]; j <= amount; j++ {
			dp[j] = dp[j] + dp[j-coins[i]]
		}
	}
	return dp[amount]
}

// test_CompletePack1 先遍历物品, 在遍历背包
func completePack(weight, value []int, bagWeight int) int {
	// 定义dp数组 和初始化
	dp := make([]int, bagWeight+1)
	// 遍历顺序
	for i := 0; i < len(weight); i++ {
		// 正序会多次添加 value[i]
		for j := weight[i]; j <= bagWeight; j++ {
			// 推导公式
			dp[j] = Max(dp[j], dp[j-weight[i]]+value[i])
			// debug
		}
	}
	return dp[bagWeight]
}

/*
dp[i][j][k] = dp[i-1][j][k]
dp[i][j][k] = dp[i-1][j-zero][k-one]
*/
func zeroOne(strs []string, m int, n int) int {

	dp := make([][][]int, len(strs)+1)
	//初始化数组
	for i, _ := range dp {
		dp[i] = make([][]int, m+1)
		for j, _ := range dp[i] {
			dp[i][j] = make([]int, n+1)
		}
	}
	//初始值
	for i := 1; i <= len(strs); i++ {
		str := strs[i-1]
		zero := strings.Count(str, "0")
		one := strings.Count(str, "1")
		for j := 0; j <= m; j++ {
			for k := 0; k <= n; k++ {
				if j < zero || k < one {
					dp[i][j][k] = dp[i-1][j][k]
				} else {
					dp[i][j][k] = Max(dp[i-1][j][k], dp[i-1][j-zero][k-one]+1)
				}
			}
		}
	}
	return dp[len(strs)][m][n]
}

/*
dp[i][j] = dp[i-1][j-nums[i]] + dp[i-1][j+nums[i]]
*/
func sumWays(nums []int, target int) int {

	sum := 0
	for i, _ := range nums {
		sum += nums[i]
	}
	if sum < target || -sum > target {
		return 0
	}
	// -sum 0 +sum
	fn := make([][]int, len(nums))

	//数组初始化
	for i, _ := range fn {
		fn[i] = make([]int, 2*sum+1)
	}
	//首行初始化
	if nums[0] == 0 {
		fn[0][sum] = 2
	} else {
		fn[0][sum-nums[0]] = 1
		fn[0][sum+nums[0]] = 1
	}

	for i := 1; i < len(nums); i++ {
		for j := 0; j < 2*sum+1; j++ {
			if j-nums[i] < 0 {
				fn[i][j] = fn[i-1][j+nums[i]]
			} else if j+nums[i] > 2*sum {
				fn[i][j] = fn[i-1][j-nums[i]]
			} else {
				fn[i][j] = fn[i-1][j-nums[i]] + fn[i-1][j+nums[i]]
			}

		}
	}
	return fn[len(nums)-1][sum+target]
}

func stoneSlpit(stones []int) int {
	sum := 0

	for i, _ := range stones {
		sum += stones[i]
	}
	target := sum / 2
	fn := make([]int, target+1)

	for i := 0; i < len(stones); i++ {
		for j := target; j >= stones[i]; j-- {
			fn[j] = Max(fn[j], fn[j-stones[i]]+stones[i])
		}
	}
	return sum - 2*fn[target]
}

func canPartitionKSubsets(nums []int, k int) bool {

	sum := 0

	maxNum := 0
	for i, _ := range nums {
		maxNum = Max(maxNum, nums[i])
		sum += nums[i]
	}
	target := sum / k

	if sum%k != 0 || maxNum > target {
		return false
	}

	return false
}

func canPartition(nums []int) bool {
	sum := 0
	maxNum := 0
	for i, _ := range nums {
		sum += nums[i]
		maxNum = Max(maxNum, nums[i])
	}

	target := sum / 2
	if sum%2 != 0 || maxNum > target {
		return false
	}

	fn := make([]bool, target+1)
	fn[0] = true

	for i := 0; i < len(nums); i++ {
		for j := target; j >= nums[i]; j-- {
			fn[j] = fn[j] || fn[j-nums[i]]
		}
	}
	return fn[target]
}

func splitNums2(nums []int) bool {
	sum := 0

	maxNum := 0
	for i, _ := range nums {
		maxNum = Max(maxNum, nums[i])
		sum += nums[i]
	}
	target := sum / 2

	if sum%2 != 0 || maxNum > target {
		return false
	}

	fmn := make([]int, target+1)

	for i := 0; i < len(nums); i++ {
		for j := target; j >= nums[i]; j-- {
			if fmn[j] < fmn[j-nums[i]]+nums[i] {
				fmn[j] = fmn[j-nums[i]] + nums[i]
			}
		}
	}
	return fmn[target] == target
}

// 给你一个 只包含正整数 的 非空 数组 nums。请你判断是否可以将这个数组分割成两个子集，使得两个子集的元素和相等。
/*
total = sum/2
重量 nums[i]  最大 sum/2
价值 nums[i]  目标 sum/2

*/
func splitNums(nums []int) bool {
	sum := 0

	for i, _ := range nums {
		sum += nums[i]
	}
	if sum%2 != 0 {
		return false
	}
	target := sum / 2

	fmn := make([][]int, len(nums))
	for i := 0; i < len(nums); i++ {
		fmn[i] = make([]int, target+1)
		fmn[i][0] = 0
	}
	for i := 0; i < target+1; i++ {
		if i >= nums[0] {
			fmn[0][i] = nums[0]
		}
	}

	for i := 1; i < len(nums); i++ {
		num := nums[i]
		for j := 0; j < target+1; j++ {
			if num > j {
				fmn[i][j] = fmn[i-1][j]
			} else {
				if fmn[i-1][j] > fmn[i-1][j-num]+num {
					fmn[i][j] = fmn[i-1][j]
				} else {
					fmn[i][j] = fmn[i-1][j-num] + num
				}
			}
		}
		if fmn[i][target] == target {
			return true
		}
	}

	return false
}

/*
01背包
*/
func bag(w []int, v []int, totalW int) int {
	if len(w) != len(v) {
		return 0
	}
	fmn := make([][]int, len(w))

	for i := 0; i < len(w); i++ {
		fmn[i] = make([]int, totalW+1)
	}

	for i := w[0]; i <= totalW; i++ {
		fmn[0][i] = v[0]
	}

	for i := 1; i < len(w); i++ {
		for j := 0; j <= totalW; j++ {
			if j < w[i] {
				fmn[i][j] = fmn[i-1][j]
			} else {
				fmn[i][j] = Max(fmn[i-1][j], fmn[i-1][j-w[i]]+v[i])
			}
		}
	}
	return fmn[len(w)-1][totalW]
}

// 滚动数组
func bag2(w []int, v []int, totalW int) int {
	if len(w) != len(v) {
		return 0
	}
	fmn := make([]int, totalW+1)

	for i := 0; i < len(w); i++ {
		for j := totalW; j >= w[i]; j-- {
			fmn[j] = Max(fmn[j], fmn[j-w[i]]+v[i])
		}
	}
	return fmn[totalW]
}

/*
二叉搜索树的个数
长度为n。思考以i为根节点，树的个数，然后i依次从0到n-1.
当1为根节点时：总个数 = 左子树个数*右子树个数 = f(0) * f(n-1)
当2为根节点时：总个数 = 左子树个数*右子树个数 = f(1) * f(n-2)
当n为根节点时：总个数 = 左子树个数*右子树个数 = f(n-1) * f(0)
所以 f(n) = f(0)*f(n-1) + f(1) * f(n-2) + ... + f(n-1) * f(0)
*/
func numTree(n int) int {
	fn := make([]int, n+1)
	fn[0] = 1
	fn[1] = 1
	for i := 2; i <= n; i++ {
		for j := 0; j < i; j++ {
			fn[i] += fn[j] * fn[i-j-1]
		}
	}
	return fn[n]
}

/*
f(i) = j * (i -j)  或者 j * f(i-j) ;j从1开始到i-1
*/
func intSlpit(n int) int {
	if n == 0 || n == 1 {
		return 0
	}
	fn := make([]int, n+1)

	for i := 2; i <= n; i++ {
		curMax := 0
		for j := 1; j < n; j++ {
			curMax = Max(curMax, Max(j*(i-j), j*fn[i-j]))
		}
		fn[i] = curMax
	}
	return fn[n]
}

// 斐波那契数 0,1,1,2
func fb(n int) (res int) {
	if n == 1 {
		return 0
	} else if n == 2 {
		return 1
	}
	a := 0
	b := 1
	fmt.Printf("%d,", a)
	fmt.Printf("%d,", b)
	for i := 2; i <= n; i++ {
		res = a + b
		a, b = b, res
		fmt.Printf("%d,", res)
	}
	return res
}

// 70. 爬楼梯  https://leetcode.cn/problems/climbing-stairs/  fn = fn-1 + fn-2
func cs(n int) (res int) {
	if n <= 2 {
		return n
	}
	fx := make([]int, n+1)
	fx[1] = 1
	fx[2] = 2
	for i := 3; i <= n; i++ {
		fx[i] = fx[i-1] + fx[i-2]
	}
	return fx[n]
}

// 使用最小花费爬楼梯：每次一步或者2步，至楼顶 [10],[10, 12],[10,15,20],[1,100,1,1,1,100,1,1,100,1]，
// [1,100,1,1,1,100,1,1,100,1]，
//
//	1,100,2,
func mincs(cost []int) (res int) {
	if len(cost) == 1 {
		return 0
	}
	fn := make([]int, len(cost))
	fn[0] = cost[0]
	fn[1] = cost[1]
	for i := 2; i < len(cost); i++ {
		fn[i] = MIN(fn[i-1], fn[i-2]) + cost[i]
	}
	return MIN(fn[len(fn)-1], fn[len(fn)-2])
}
func MIN(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

/*
	m

n
*/
func diffPaths(m int, n int) int {
	fmn := make([][]int, m)
	for i := 0; i < m; i++ {
		temp := make([]int, n)
		temp[0] = 1
		fmn[i] = temp
	}

	for i := 0; i < n; i++ {
		fmn[0][i] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			fmn[i][j] = fmn[i][j-1] + fmn[i-1][j]
		}
	}

	return fmn[m-1][n-1]
}

func diffPaths01(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])
	fmn := make([][]int, m)
	for i := 0; i < m; i++ {
		temp := make([]int, n)
		fmn[i] = temp
	}

	for i := 0; i < m && obstacleGrid[i][0] == 0; i++ {
		//如果有障碍物 ==1 ，后续都不可达。
		fmn[i][0] = 1
	}

	for i := 0; i < n && obstacleGrid[0][i] == 0; i++ {
		//如果 ==1 后续不可达
		fmn[0][i] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] == 0 {
				fmn[i][j] = fmn[i][j-1] + fmn[i-1][j]
			}
		}
	}
	return fmn[m-1][n-1]
}
