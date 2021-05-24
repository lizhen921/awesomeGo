package reverselist

/*
反转一个单链表。

示例:

输入: 1->2->3->4->5->NULL
输出: 5->4->3->2->1->NULL
进阶:
你可以迭代或递归地反转链表。你能否用两种方法解决这道题？

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/reverse-linked-list

*/
import "fmt"

//反转整个链表 -- 迭代
func reverseList(head *ListNode) *ListNode {
	cur := head
	var pre *ListNode
	for cur != nil {
		pre, cur, cur.Next = cur, cur.Next, pre
	}
	return pre
}

//反转整个链表 -- 递归
func reverseList2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := reverseList2(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}

func reverseNode(preNode *ListNode) (afterNode *ListNode) {
	if preNode.Next == nil {
		return afterNode
	}
	return reverseNode(preNode)
}

//K链表反转，k个为一组，不足不反转
func reverseKList(head *ListNode, k int) *ListNode {
	hair := &ListNode{
		Next: head, Value: -1,
	}
	pre := hair
	tail := hair.Next
	var next *ListNode
	i := 1
	for tail != nil {
		if i%k == 0 {
			//记录下一个链表的
			next = tail.Next

			//暂时断开要反转的链表与后面元素的链接
			tail.Next = nil

			//反转head链表
			childHeadNew := reverseList(head)

			//由于原链表反转，tail跟着成为第一个元素，head成了最后一个，需要找到正确的tail值
			tail = head

			//前置节点 重新链接 反转后的链表
			pre.Next = childHeadNew

			//重新链接之前断开的链表
			tail.Next = next

			//重新定义pre节点和head链表
			pre = tail
			head = next
		}
		tail = tail.Next
		i++
	}

	return hair.Next
}

func KlistReverse(head *ListNode, k int) *ListNode {
	hari := &ListNode{
		Next: head,
	}

	pre := hari
	tail := head
	var nextHead *ListNode

	i := 1
	for tail != nil {
		if i%k == 0 {
			//记录下个链表表头
			nextHead = tail.Next
			//切断链表
			tail.Next = nil

			//反转子链表
			tail = reverseList(head)

			//重新赋值变量
			tail, head = head, tail

			//链接链表
			tail.Next = nextHead
			pre.Next = head

			//重新赋值变量
			pre = tail
			head = nextHead

		}
		i++
		tail = tail.Next
	}

	return hari.Next
}

//打印链表
func PrintNodes(list *ListNode) {
	for list != nil {
		fmt.Printf("%d", list.Value)
		list = list.Next
	}
	fmt.Println("")
}
