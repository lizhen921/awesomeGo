package twosum

import (
	"fmt"
	"testing"
)

func TestTwoSum(t *testing.T) {

	l1 := new(ListNode)
	l1.Val = 5
	//l1.Next = new(ListNode)
	//l1.Next.Val = 9
	//l1.Next.Next = new(ListNode)
	//l1.Next.Next.Val = 1
	//
	l2 := new(ListNode)
	l2.Val = 5
	//cur2 := l2
	//for i := 1; i < 4; i++ {
	//	next := new(ListNode)
	//	next.Val = 9
	//	cur2.Next = next
	//	cur2 = next
	//}

	listNode := addTwoNumbers(l1, l2)
	PrintNodes(listNode)

}

//打印链表
func PrintNodes(list *ListNode) {
	for list != nil {
		fmt.Printf("%d", list.Val)
		list = list.Next
	}
	fmt.Println("")
}
