package reverselist

/*
给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。

进阶：你能尝试使用一趟扫描实现吗？

*/

//借用栈的方式
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head.Next == nil && n == 1 {
		return nil
	}
	root := head
	nodeChan := make(chan *ListNode, n+1)

	for root != nil {
		if root.Next == nil && len(nodeChan) < n {
			return head.Next
		}
		nodeChan <- root

		if len(nodeChan) == n+1 {

			tempNode := <-nodeChan

			if root.Next == nil {
				tempNode.Next = tempNode.Next.Next
			}
		}
		root = root.Next
	}
	return head
}

/**
2，3，4 都是一次遍历，且使用O(1)的空间，不断优化的代码
*/
func removeNthFromEnd2(head *ListNode, n int) *ListNode {
	if head == nil {
		return head
	}

	pre := new(ListNode)
	pre.Next = head
	first := pre
	second := pre

	distance := 0
	for first.Next != nil {
		if distance >= n {
			second = second.Next
		}

		first = first.Next
		distance++
	}
	if distance >= n {
		if second.Next == head {
			return head.Next
		}
		second.Next = second.Next.Next
	}
	return head
}

func removeNthFromEnd3(head *ListNode, n int) *ListNode {
	if head == nil {
		return head
	}
	pre := new(ListNode)
	pre.Next = head

	first := head
	second := pre

	for i := 0; i < n; i++ {
		first = first.Next
	}

	for ; first != nil; first = first.Next {
		second = second.Next
	}

	second.Next = second.Next.Next
	return pre.Next
}

func removeNthFromEnd4(head *ListNode, n int) *ListNode {
	dummy := &ListNode{head, 0}
	first, second := head, dummy
	for i := 0; i < n; i++ {
		first = first.Next
	}
	for ; first != nil; first = first.Next {
		second = second.Next
	}
	second.Next = second.Next.Next
	return dummy.Next
}
