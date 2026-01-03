package main

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var carry, total int
	dummy := &ListNode{Val: -1}
	curr := dummy
	for l1 != nil || l2 != nil {
		total = carry
		if l1 != nil {
			total += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			total += l2.Val
			l2 = l2.Next
		}
		curr.Next = &ListNode{Val: total % 10}
		carry = total / 10
		curr = curr.Next
	}
	if carry != 0 {
		curr.Next = &ListNode{Val: carry}
	}
	return dummy.Next
}
