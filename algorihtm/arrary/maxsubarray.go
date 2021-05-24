package arrary

/**
给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。

示例:

输入: [-2,1,-3,4,-1,2,1,-5,4]
输出: 6
解释: 连续子数组 [4,-1,2,1] 的和最大，为 6。
进阶:

如果你已经实现复杂度为 O(n) 的解法，尝试使用更为精妙的分治法求解。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/maximum-subarray
*/

/**
分治法：
 假设【以i个数为结尾】的【连续子数组最大和】是f(i), 那么如何求f(i)？ 假设知道f(i-1)，以及i位置的数nums[i], 那么得到结论：
	f(i) = max{f(i-1)+nums[i], nums[i]}
 根据这个公式不难计算出改数组中的最大和是什么
*/

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	max := nums[0]
	maxi := 0 //第i个元素为结尾的最大和
	for i, _ := range nums {
		if i == 0 {
			maxi = nums[0]
		} else {
			if maxi+nums[i] < nums[i] {
				maxi = nums[i]
			} else {
				maxi = maxi + nums[i]
			}
		}
		if maxi > max {
			max = maxi
		}
	}
	return max
}
