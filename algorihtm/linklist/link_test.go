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
	if this.Len == 0{
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

	if index <0 || index >= this.Len {
		return
	} else if index == 0 {
		this.Head = this.Head.Next
	} else {
		node := this.Head
		for i := 0; i < index-1; i++ {
			node = node.Next
		}
		node.Next = node.Next.Next
		if index == this.Len - 1 {
			this.Tail = node
		}
	}
	this.Len--
}

func TestLinkList(t *testing.T) {
	mylist := Constructor()
	mylist.AddAtHead(0)
	mylist.AddAtIndex(1,4)
	mylist.AddAtTail(8)
	mylist.AddAtHead(5)
	mylist.AddAtIndex(4,3)
	mylist.AddAtTail(0)
	mylist.AddAtTail(5)
	mylist.AddAtIndex(6,3)
	mylist.DeleteAtIndex(7)
	mylist.DeleteAtIndex(5)
	mylist.AddAtTail(4)
}
