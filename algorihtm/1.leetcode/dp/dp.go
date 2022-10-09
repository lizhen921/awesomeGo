package dp

//0、1、1、2、3、5...
func fib(n int) []int {
	dp := make([]int, 0)
	for i := 0; i <= n; i++ {
		if i > 1 {
			temp := dp[i-1] + dp[i-2]
			dp = append(dp, temp)
		} else if i == 0 {
			dp = append(dp, i)
		} else if i == 1 {
			dp = append(dp, i)
		}
	}
	return dp
}

func climbStairs01(n int) []int {

	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		if i == 1 {
			dp[i] = 1
		} else if i == 2 {
			dp[i] = 2
		} else {
			dp[i] = dp[i-1] + dp[i-2]
		}
	}
	return dp
}

func climbStairs02(cost []int) int {
	if len(cost) == 1 {
		return 0
	}
	if len(cost) == 2 {
		return min(cost[0], cost[1])
	}
	a := cost[0]
	b := cost[1]

	for i := 2; i < len(cost); i++ {
		c := min(a+cost[i], b+cost[i])
		a = b
		b = c
	}

	return min(a, b)
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func uniquePaths(m int, n int) int {

	dp := make([][]int, m)

	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		dp[i][0] = 1
	}
	for i := 0; i < n; i++ {
		dp[0][i] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[m-1][n-1]
}

func uniquePathsWithObstacles(obstacleGrid [][]int) int {

	m := len(obstacleGrid)

	n := 0
	if m > 0 {
		n = len(obstacleGrid[0])
	}

	dp := make([][]int, m)

	isIPath := false
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		if obstacleGrid[i][0] == 1 || isIPath {
			dp[i][0] = 0
			isIPath = true
		} else {
			dp[i][0] = 1
		}
	}
	isIPath = false
	for i := 0; i < n; i++ {
		if obstacleGrid[0][i] == 1 || isIPath {
			dp[0][i] = 0
			isIPath = true
		} else {
			dp[0][i] = 1
		}
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] == 1 {
				dp[i][j] = 0
			} else {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}
	return dp[m-1][n-1]
}

func integerBreak(n int) int {

}
