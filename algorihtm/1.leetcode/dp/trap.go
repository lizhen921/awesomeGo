package dp

/**

给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。

示例 1：

输入：height = [0,1,0,2,1,0,1,3,2,1,2,1]
输出：6
解释：上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/trapping-rain-water
*/

/*
一个位置能存放水的个数，取决于它两侧最大高度的较小值减去当前位置高度

*/
func trap(height []int) int {

	return 0
}

/*
双指针，两个指针分别从首尾开始，表示最大值，逐个遍历各个位置，同时移动指针，更新最大值
*/
func trap2(height []int) int {
	if len(height) <= 0 {
		return 0
	}
	left, left_max := 0, height[0]
	right, right_max := len(height)-1, height[len(height)-1]

	count := 0
	for i := 0; i < len(height); i++ {
		curHeight := 0
		if left_max < right_max {
			curHeight = height[left]
			if curHeight < left_max {
				count = count + (left_max - curHeight)
			} else {
				left_max = curHeight
			}
			left++
		} else {
			curHeight = height[right]
			if curHeight < right_max {
				count = count + (right_max - curHeight)
			} else {
				right_max = curHeight
			}
			right--
		}
	}

	return count
}
