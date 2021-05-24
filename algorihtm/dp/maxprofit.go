package dp

/**

给定一个数组 prices ，它的第 i 个元素 prices[i] 表示一支给定股票第 i 天的价格。

你只能选择 某一天 买入这只股票，并选择在 未来的某一个不同的日子 卖出该股票。设计一个算法来计算你所能获取的最大利润。

返回你可以从这笔交易中获取的最大利润。如果你不能获取任何利润，返回 0 。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock

*/
/*
动态规划：
 假设第i天卖出股票，则需要找到第i天之前第最小值min(i), 来计算第i天卖出的话，最大收益是多少。依次计算（0—n）天，找出哪天卖出有最大值即可。

*/
func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	maxPro := 0
	minNum := prices[0]
	for _, v := range prices {
		if minNum > v {
			minNum = v
		}
		if v-minNum > maxPro {
			maxPro = v - minNum
		}

	}
	return maxPro
}

func maxProfit2(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	minNum, maxPro := prices[0], 0
	for i := 0; i < len(prices); i++ {
		if minNum > prices[i] {
			minNum = prices[i]
		} else if prices[i]-minNum > maxPro {
			maxPro = prices[i] - minNum
		}

	}
	return maxPro
}
