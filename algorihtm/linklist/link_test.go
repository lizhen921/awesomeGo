package linklist

import (
	"fmt"
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
	Pre  *ListNode
}

func GengrateLinkList(valList []int) *ListNode {
	var head *ListNode
	var tail *ListNode

	for i := 0; i < len(valList); i++ {
		node := &ListNode{
			Val:  valList[i],
			Next: nil,
		}
		if tail == nil {
			tail = node
			head = tail
		} else {
			tail.Next = node
			tail = node
		}
	}
	return head
}

func PrintLinkList(head *ListNode) {
	for head != nil {
		fmt.Printf("%d->", head.Val)
		head = head.Next
	}
	fmt.Println("")
}

func TestGengrateLinkList(t *testing.T) {
	valList := []int{3, 4, 6, 1, 7, 2}
	head := GengrateLinkList(valList)
	PrintLinkList(head)
}

/*
给你一个链表的头节点 head 和一个整数 val ，请你删除链表中所有满足 Node.val == val 的节点，并返回新的头节点 。

输入：head = [1,2,6,3,4,5,6], val = 6
输出：[1,2,3,4,5]
*/
func removeElements(head *ListNode, val int) *ListNode {
	preNode := &ListNode{
		Next: head,
	}
	for tempHead := preNode; tempHead.Next != nil; {
		if tempHead.Next.Val == val {
			tempHead.Next = tempHead.Next.Next
		} else {
			tempHead = tempHead.Next
		}
	}
	return preNode.Next
}

func TestRemoveElements(t *testing.T) {
	valList := []int{7, 7, 7, 8}
	head := GengrateLinkList(valList)
	PrintLinkList(head)
	head = removeElements(head, 7)
	PrintLinkList(head)
}

type MyLinkedList struct {
	Head *ListNode
	Tail *ListNode
	Len  int
}

func Constructor() MyLinkedList {
	return MyLinkedList{
		Len: 0,
	}
}

func (this *MyLinkedList) Get(index int) int {
	if index >= this.Len {
		return -1
	}
	node := this.Head
	for i := 0; i < index; i++ {
		node = node.Next
	}
	return node.Val
}

func (this *MyLinkedList) AddAtHead(val int) {
	head := &ListNode{
		Val: val,
	}
	head.Next = this.Head
	this.Head = head
	if this.Len == 0 {
		this.Tail = this.Head
	}
	this.Len++
}

func (this *MyLinkedList) AddAtTail(val int) {
	tail := &ListNode{
		Val: val,
	}
	if this.Tail == nil {
		this.Tail = tail
		this.Head = tail
	} else {
		this.Tail.Next = tail
		this.Tail = tail
	}
	this.Len++
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	addNode := &ListNode{
		Val: val,
	}
	node := this.Head
	if index <= 0 {
		this.AddAtHead(val)
	} else if index == this.Len {
		this.AddAtTail(val)
	} else if index > this.Len {
		return
	} else {
		for i := 0; i < index-1; i++ {
			node = node.Next
		}
		temp := node.Next
		node.Next = addNode
		addNode.Next = temp
		this.Len++
	}

}

func (this *MyLinkedList) DeleteAtIndex(index int) {

	if index < 0 || index >= this.Len {
		return
	} else if index == 0 {
		this.Head = this.Head.Next
	} else {
		node := this.Head
		for i := 0; i < index-1; i++ {
			node = node.Next
		}
		node.Next = node.Next.Next
		if index == this.Len-1 {
			this.Tail = node
		}
	}
	this.Len--
}

func TestLinkList(t *testing.T) {
	mylist := Constructor()
	mylist.AddAtHead(0)
	mylist.AddAtIndex(1, 4)
	mylist.AddAtTail(8)
	mylist.AddAtHead(5)
	mylist.AddAtIndex(4, 3)
	mylist.AddAtTail(0)
	mylist.AddAtTail(5)
	mylist.AddAtIndex(6, 3)
	mylist.DeleteAtIndex(7)
	mylist.DeleteAtIndex(5)
	mylist.AddAtTail(4)
}

/*
反转链表：
给你单链表的头节点 head ，请你反转链表，并返回反转后的链表。

1 --> 2 --> 3 --> 4 --> 5

5 --> 4 --> 3 --> 2 --> 1
*/

func reverseList(head *ListNode) *ListNode {

	var pre *ListNode
	cur := head
	for cur.Next != nil {
		curNext := cur.Next
		cur.Next = pre
		pre = cur
		cur = curNext
	}
	cur.Next = pre

	return cur
}

func reverseList2(head *ListNode) *ListNode {
	var pre *ListNode
	cur := head

	for cur.Next != nil {
		cur.Next, pre, cur = pre, cur, cur.Next
	}
	cur.Next = pre
	return cur
}

func TestReverseList(t *testing.T) {
	valList := []int{3, 4, 6, 1, 7, 2}
	head := GengrateLinkList(valList)
	PrintLinkList(head)
	head = reverseList3(head)
	PrintLinkList(head)

}

func reverseList3(head *ListNode) *ListNode {

	if head == nil {
		return head
	}

	var pre *ListNode
	cur := head

	for cur.Next != nil {
		temp := cur.Next
		cur.Next = pre
		pre = cur
		cur = temp
	}
	cur.Next = pre

	return cur
}

/*
两两交换链表元素

给定一个链表，两两交换其中相邻的节点，并返回交换后的链表。

你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。
*/

func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{
		Next: head,
	}

	tail := dummy

	pre := dummy
	cur := head
	swap := false
	for cur != nil {
		if swap {
			temp := cur.Next
			cur.Next = pre
			pre.Next = temp
			tail.Next = cur
			tail = pre
			cur = tail.Next
			swap = false
		} else {
			pre = pre.Next
			cur = cur.Next
			swap = true
		}
	}
	return dummy.Next
}

func swappairs2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	tail := head.Next
	head.Next = swappairs2(tail.Next)
	tail.Next = head

	return tail
}

func TestSwapPairs(t *testing.T) {
	valList := []int{3, 4, 6, 1, 7, 2, 1}
	head := GengrateLinkList(valList)
	PrintLinkList(head)
	head = swappairs2(head)
	PrintLinkList(head)
}

/*
存在一个按升序排列的链表，给你这个链表的头节点 head ，请你删除所有重复的元素，使每个元素 只出现一次 。

返回同样按升序排列的结果链表。

输入：head = [1,1,2]
输出：[1,2]
*/
func deleteDuplicates(head *ListNode) *ListNode {

	if head == nil || head.Next == nil {
		return head
	}

	head.Next = deleteDuplicates(head.Next)

	if head.Val == head.Next.Val {
		head = head.Next
	}

	return head
}

/*
给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。
[1,2]
1
*/

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	head, index := remove(head, n)
	if index == n {
		head = head.Next
	}
	return head
}

func remove(head *ListNode, n int) (node *ListNode, index int) {
	if head == nil {
		return head, 0
	}

	next, index := remove(head.Next, n)

	if n == index && next != nil {
		head.Next = next.Next
	}

	return head, index + 1
}

func TestRemove(t *testing.T) {
	valList := []int{1, 2, 3, 4, 5, 6, 7, 8}
	head := GengrateLinkList(valList)
	PrintLinkList(head)
	//递归
	head = removeNthFromEnd(head, 1)
	PrintLinkList(head)
	head = removeNthFromEnd2(head, 4)
	PrintLinkList(head)
}

func removeNthFromEnd2(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: head}
	first := dummy
	second := dummy
	for i := 0; i < n; i++ {
		first = first.Next
	}
	for first.Next != nil {
		first = first.Next
		second = second.Next
	}
	if second != nil && second.Next != nil {
		second.Next = second.Next.Next
	}
	return dummy.Next
}

/*
























 */
