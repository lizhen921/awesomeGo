package stack

import "testing"

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
