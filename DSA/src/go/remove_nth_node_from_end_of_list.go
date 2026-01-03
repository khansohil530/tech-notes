package main

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Val: 0, Next: head}
	first, second := dummy, dummy
	for n > 0 && second != nil {
		second = second.Next
		n--
	}
	for second != nil && second.Next != nil {
		first = first.Next
		second = second.Next
	}
	first.Next = first.Next.Next
	return dummy.Next
}
