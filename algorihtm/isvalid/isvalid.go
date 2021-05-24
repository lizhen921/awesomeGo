package isvalid

/*
给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。

有效字符串需满足：

左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。
注意空字符串可被认为是有效字符串。
链接：https://leetcode-cn.com/problems/valid-parentheses
*/
func isValid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}
	var validMap = map[byte]byte{
		')': '(',
		'}': '{',
		']': '[',
	}

	var listTemp []byte
	for i, _ := range s {
		if v, ok := validMap[s[i]]; ok {
			if len(listTemp) > 0 && listTemp[len(listTemp)-1] == v {
				listTemp = listTemp[:len(listTemp)-1]
			} else {
				return false
			}

		} else {
			listTemp = append(listTemp, s[i])
		}
	}
	return len(listTemp) == 0
}
