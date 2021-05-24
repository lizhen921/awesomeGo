package reverselist

import "testing"

func TestRemove(t *testing.T) {
	head := new(ListNode)
	cur := head
	for i := 0; i < 1; i++ {
		next := new(ListNode)
		next.Value = i + 1
		cur.Next = next
		cur = next
	}

	PrintNodes(head)
	root := removeNthFromEnd3(head, 2)
	PrintNodes(root)
}
