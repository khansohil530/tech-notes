package main

func reorderList(head *ListNode) {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	curr := slow.Next
	slow.Next = nil
	var prev, next *ListNode

	for curr != nil {
		next = curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	first, second := head, prev
	var next1, next2 *ListNode
	for first != nil && second != nil {
		next1, next2 = first.Next, second.Next
		first.Next = second
		second.Next = next1
		first, second = next1, next2
	}
}
