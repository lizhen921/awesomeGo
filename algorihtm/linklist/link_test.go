package linklist

import (
	"fmt"
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
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