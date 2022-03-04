package __jianzhioffer

type ListNode struct {
	Next *ListNode
	Val  int
}

//递归方式
func reversePrint(head *ListNode) (res []int) {
	var f func(*ListNode)
	f = func(h *ListNode) {
		if h == nil {
			return
		}
		f(h.Next)
		res = append(res, h.Val)
	}

	f(head)
	return res
}

//利用栈
func reversePrint02(head *ListNode) (res []int) {
	stack := make([]int, 0)

	for head != nil {
		stack = append(stack, head.Val)
		head = head.Next
	}
	for i := len(stack) - 1; i >= 0; i-- {
		res = append(res, stack[i])
	}
	return res
}

//反转链表，顺序遍历
func reversePrint03(head *ListNode) (res []int) {
	head = reverseLink(head)
	for head != nil {
		res = append(res, head.Val)
	}
	return res
}

//反转链表方法1
func reverseLink(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := reverseLink(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}

//反转链表方法2
func reverseLink02(head *ListNode) *ListNode {
	var pre *ListNode
	cru := head

	for cru != nil {
		temp := cru.Next
		cru.Next = pre
		pre = cru
		cru = temp
	}
	return pre
}
