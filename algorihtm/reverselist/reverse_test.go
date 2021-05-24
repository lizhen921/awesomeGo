package reverselist

import "testing"

//反转链表测试 --迭代
func TestReverseList(t *testing.T) {
	//初始化 链表
	head := new(ListNode)
	cur := head
	for i := 0; i < 9; i++ {
		next := new(ListNode)
		next.Value = i + 1
		cur.Next = next
		cur = next
	}

	PrintNodes(head)
	revList := reverseList(head)
	PrintNodes(revList)
}

//反转链表测试 -- 递归
func TestReverseList2(t *testing.T) {
	//初始化 链表
	head := new(ListNode)
	cur := head
	for i := 0; i < 9; i++ {
		next := new(ListNode)
		next.Value = i + 1
		cur.Next = next
		cur = next
	}

	PrintNodes(head)
	revList := reverseList2(head)
	PrintNodes(revList)
}

//K链表反转，k个为一组，不足不反转
func TestReverseKList(t *testing.T) {
	//初始化 链表
	head := new(ListNode)
	cur := head
	for i := 0; i < 9; i++ {
		next := new(ListNode)
		next.Value = i + 1
		cur.Next = next
		cur = next
	}
	PrintNodes(head)
	revKList := reverseKList(head, 4)
	PrintNodes(revKList)
}

//K链表反转，k个为一组，不足不反转
func TestReverseKkList(t *testing.T) {
	//初始化 链表
	head := new(ListNode)
	cur := head
	for i := 0; i < 9; i++ {
		next := new(ListNode)
		next.Value = i + 1
		cur.Next = next
		cur = next
	}
	PrintNodes(head)
	revKList := KlistReverse(head, 4)
	PrintNodes(revKList)
}
