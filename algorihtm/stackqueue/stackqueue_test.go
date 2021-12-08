package stack

import (
	"strconv"
	"testing"
)

type MyStack struct {
	queue []int //创建一个队列
}

/** Initialize your data structure here. */
func Constructor() MyStack {
	return MyStack{ //初始化
		queue: make([]int, 0),
	}
}

/** Push element x onto stack. */
func (this *MyStack) Push(x int) {
	//添加元素
	this.queue = append(this.queue, x)
}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	//n:=len(this.queue)-1//判断长度
	//for n!=0{ //除了最后一个，其余的都重新添加到队列里
	//	val:=this.queue[0]
	//	this.queue=this.queue[1:]
	//	this.queue=append(this.queue,val)
	//	n--
	//}
	////弹出元素
	//val:=this.queue[0]
	//this.queue=this.queue[1:]
	//return val

	n := len(this.queue)

	val := this.queue[n-1]
	this.queue = this.queue[:n-1]
	return val

}

/** Get the top element. */
func (this *MyStack) Top() int {
	//利用Pop函数，弹出来的元素重新添加
	val := this.Pop()
	this.queue = append(this.queue, val)
	return val
}

/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	return len(this.queue) == 0
}

/*
有效的括号
给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。

有效字符串需满足：

左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。
{[]}

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

func TestIsValid(t *testing.T) {
	isValid("]")
}

/*
删除字符串中的所有相邻重复项
给出由小写字母组成的字符串 S，重复项删除操作会选择两个相邻且相同的字母，并删除它们。

在 S 上反复执行重复项删除操作，直到无法继续删除。

在完成所有重复项删除操作后返回最终的字符串。答案保证唯一。

示例：

输入："abbaca"
输出："ca"
解释：
例如，在 "abbaca" 中，我们可以删除 "bb" 由于两字母相邻且相同，这是此时唯一可以执行删除操作的重复项。
之后我们得到字符串 "aaca"，其中又只有 "aa" 可以执行重复项删除操作，所以最后的字符串为 "ca"。

*/
func removeDuplicates(s string) string {
	if len(s) <= 1 {
		return s
	}
	strStack := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		if len(strStack) > 0 && s[i] == strStack[len(strStack)-1] {
			strStack = strStack[:len(strStack)-1]
			continue
		}
		strStack = append(strStack, s[i])
	}
	return string(strStack)
}

func TestRemoveDu(t *testing.T) {
	removeDuplicates("abbaca")
}

/*
逆波兰表达式求值

根据 逆波兰表示法，求表达式的值。

有效的算符包括 +、-、*、/ 。每个运算对象可以是整数，也可以是另一个逆波兰表达式。


说明：

整数除法只保留整数部分。
给定逆波兰表达式总是有效的。换句话说，表达式总会得出有效数值且不存在除数为 0 的情况。

示例 1：

输入：tokens = ["2","1","+","3","*"]
输出：9
解释：该算式转化为常见的中缀算术表达式为：((2 + 1) * 3) = 9

链接：https://leetcode-cn.com/problems/evaluate-reverse-polish-notation

*/
func evalRPN(tokens []string) int {
	res := 0
	if len(tokens) == 1 {
		res, _ = strconv.Atoi(tokens[0])
		return res
	}
	kMap := map[string]bool{"+": true, "-": true, "*": true, "/": true}

	tempStack := make([]string, 0)
	for _, v := range tokens {
		if _, ok := kMap[v]; ok {
			tempValueA, _ := strconv.Atoi(tempStack[len(tempStack)-2])
			tempValueB, _ := strconv.Atoi(tempStack[len(tempStack)-1])
			tempAB := 0
			switch v {
			case "+":
				tempAB = tempValueA + tempValueB
			case "-":
				tempAB = tempValueA - tempValueB
			case "*":
				tempAB = tempValueA * tempValueB
			case "/":
				tempAB = tempValueA / tempValueB
			default:
				return -1
			}
			tempStack = tempStack[:len(tempStack)-2]
			tempStack = append(tempStack, strconv.Itoa(tempAB))
			res = tempAB
		} else {
			tempStack = append(tempStack, v)
		}
	}
	return res
}

func TestEvalRpn(t *testing.T) {
	evalRPN([]string{"4", "13", "5", "/", "+"})
}

/*
滑动窗口最大值
*/

func maxSlidingWindow(nums []int, k int) []int {
	l := len(nums)
	if l <= 1 || k == 1 {
		return nums
	}
	maxIndex := make([]int, 0)

	res := make([]int, 0)
	j := 0

	for i := 0; i < l; i++ {
		if len(maxIndex) == 0 || nums[maxIndex[len(maxIndex)-1]] > nums[i] {
			maxIndex = append(maxIndex, i)
		} else {
			for len(maxIndex) > 0 && nums[maxIndex[len(maxIndex)-1]] < nums[i] {
				maxIndex = maxIndex[:len(maxIndex)-1]
			}
			maxIndex = append(maxIndex, i)
		}
		if i+1 >= k {
			if len(maxIndex) > 0 {
				res = append(res, nums[maxIndex[0]])
				if maxIndex[0] == j {
					maxIndex = maxIndex[1:]
				}
			}
			j++
		}

	}
	return res
}

func TestMax(t *testing.T) {
	nums := []int{-6, -10, -7, -1, -9, 9, -8, -4, 10, -5, 2, 9, 0, -7, 7, 4, -2, -10, 8, 7}
	maxSlidingWindow(nums, 7)
}
