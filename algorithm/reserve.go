package algorithm

type ListNode struct {
	Val  int
	Next *ListNode
}

func Reserve(head *ListNode) *ListNode {
	if head.Next == nil || head == nil {
		return nil
	}
	newHead := Reserve(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}
