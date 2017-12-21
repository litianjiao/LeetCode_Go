package lc0083

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {

	if head == nil {
		return head
	}
	pre, current := head, head
	for current != nil {
		for current.Next != nil && current.Val == current.Next.Val {
			current = current.Next
		}
		pre.Next = current.Next
		pre, current = pre.Next, current.Next
	}
	return head
}
