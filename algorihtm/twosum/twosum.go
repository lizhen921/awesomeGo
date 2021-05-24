package twosum

/*
给你两个 非空 的链表，表示两个非负的整数。它们每位数字都是按照 逆序 的方式存储的，并且每个节点只能存储 一位 数字。

请你将两个数相加，并以相同形式返回一个表示和的链表。

你可以假设除了数字 0 之外，这两个数都不会以 0 开头。
实例 1：
输入：l1 = [2,4,3], l2 = [5,6,4]
输出：[7,0,8]
解释：342 + 465 = 807.

示例 2：
输入：l1 = [0], l2 = [0]
输出：[0]

示例 3：
输入：l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
输出：[8,9,9,9,0,0,0,1]

提示：

每个链表中的节点数在范围 [1, 100] 内
0 <= Node.val <= 9
题目数据保证列表表示的数字不含前导零

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/add-two-numbers
*/
type ListNode struct {
	Val  int
	Next *ListNode
}

//https://leetcode-cn.com/problems/add-two-numbers/
//
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	needAdd := 0
	head1 := l1

	for l1 != nil && l2 != nil {
		sum := l1.Val + l2.Val + needAdd
		needAdd = 0
		if sum >= 10 {
			needAdd = 1
			sum = sum - 10
		}
		l1.Val, l2.Val = sum, sum

		if l1.Next == nil && l2.Next == nil && needAdd > 0 {
			l1.Next = new(ListNode)
			l1 = l1.Next
			break
		}

		if l1.Next == nil {
			l1.Next = l2.Next
			l1 = l1.Next
			break
		}
		if l2.Next == nil {
			l2.Next = l1.Next
			l1 = l1.Next
			break
		}

		l1 = l1.Next
		l2 = l2.Next
	}

	for l1 != nil {
		sum := l1.Val + needAdd
		needAdd = 0
		if sum >= 10 {
			needAdd = 1
			sum = sum - 10
		}
		l1.Val = sum

		if l1.Next == nil && needAdd > 0 {
			l1.Next = new(ListNode)
			l1.Next.Val = needAdd
			break
		}
		l1 = l1.Next

	}

	return head1
}
