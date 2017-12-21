package lc0083

type ListNode struct {
	val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {

	if head == nil {
		return head
	}
	pre, current := head, head
	for current != nil {
		for current.Next != nil && current.val == current.Next.val {
			current = current.Next
		}
		pre.Next = current.Next
		pre, current = pre.Next, current.Next
	}
	return head
}
