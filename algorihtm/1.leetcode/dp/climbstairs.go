package dp

/**
假设你正在爬楼梯。需要 n 阶你才能到达楼顶。

每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？

注意：给定 n 是一个正整数。

https://leetcode-cn.com/problems/climbing-stairs/
*/

/*

f(n) = f(n-1) + f(n-2)


*/
func climbStairs(n int) int {
	if n == 0 {
		return 1
	}
	brePre := 0
	pre := 1
	count := 0
	step := 1

	for step <= n {
		count = brePre + pre
		brePre = pre
		pre = count
		step++
	}

	return count
}
