package main

func getKthNode(node *ListNode, k int) *ListNode {
	for node != nil && k > 0 {
		node = node.Next
		k--
	}
	return node
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	dummy := &ListNode{Val: -1, Next: head}
	group_prev := dummy
	var kthNode, group_next, prev, curr, temp *ListNode
	for true {
		kthNode = getKthNode(group_prev, k)
		if kthNode == nil {
			break
		}

		group_next = kthNode.Next
		prev, curr = group_next, group_prev.Next
		for curr != group_next {
			temp = curr.Next
			curr.Next = prev
			prev = curr
			curr = temp
		}

		temp = group_prev.Next
		group_prev.Next = kthNode
		group_prev = temp
	}
	return dummy.Next
}
