package substring

/*
给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。
示例 1:
输入: s = "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
示例 2:

输入: s = "bbbbb"
输出: 1
解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
示例 3:

输入: s = "pwwkew"
输出: 3
解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
示例 4:

输入: s = ""
输出: 0
提示：
0 <= s.length <= 5 * 104
s 由英文字母、数字、符号和空格组成
来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/longest-substring-without-repeating-characters
*/

//解题思路：滑动窗口。出现重复数据，移除左边的就可以
func lengthOfLongestSubstring(testString string) int {
	curMap := make(map[uint8]int)
	maxLen := 0
	start := -1

	for i := 0; i < len(testString); i++ {
		v := testString[i]
		if index, ok := curMap[v]; ok {
			for start < index {
				start++
				delete(curMap, testString[start])
			}

		}
		curMap[v] = i

		curLen := i - start
		if curLen > maxLen {
			maxLen = curLen
		}
	}

	return maxLen
}
