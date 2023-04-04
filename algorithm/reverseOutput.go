package algorithm

import "fmt"

func ReverseOutput(input []int) {
	var head *ListNode
	var ptr *ListNode
	if len(input) == 0 {
		return
	}
	head = &ListNode{Val: input[0]}
	ptr = head
	for i := 1; i < len(input); i++ {
		ptr.Next = &ListNode{Val: input[i]}
		ptr = ptr.Next
	}

	reverseList(head)
}

func reverseList(head *ListNode) {
	if head == nil {
		return
	}
	reverseList(head.Next)
	fmt.Println(head.Val)
}
