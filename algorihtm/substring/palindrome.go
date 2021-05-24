package substring

/*
给你一个字符串 s，找到 s 中最长的回文子串。

示例 1：

输入：s = "babad"
输出："bab"
解释："aba" 同样是符合题意的答案。
示例 2：

输入：s = "cbbd"
输出："bb"
示例 3：

输入：s = "a"
输出："a"
示例 4：

输入：s = "ac"
输出："a"


提示：

1 <= s.length <= 1000
s 仅由数字和英文字母（大写和/或小写）组成

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/longest-palindromic-substring
*/

/*
	中心扩算:
	1. 如果回文串是奇数串，则中心字符左右两边的字符相同
	2. 如果回文串是偶数串，则中心的两个字符肯定是相同的
	3. 所以可以将相邻且相同的当作中心字符
*/
func longestPalindrome(s string) string {
	size := len(s)
	index := 0 //起始中心
	var maxLenTemp int
	var left int
	var right int
	var finalL int
	var finalR int
	for index < size { //以index位置字符为中心，逐渐右移，找到最大回文字串
		left = index
		right = index
		//找左边相同的字符
		for left > 0 && s[left-1] == s[index] {
			left--
		}
		//找右边相同的字符
		for right < size-1 && s[right+1] == s[index] {
			right++
		}
		index = right + 1
		//处理左右两边相同字符
		for left > 0 && right < size-1 && s[left-1] == s[right+1] {
			left--
			right++
		}

		if right+1-left > maxLenTemp {
			finalL = left
			finalR = right
			maxLenTemp = right + 1 - left
		}
	}

	return s[finalL : finalR+1]
}

/*
动态规划:


*/
