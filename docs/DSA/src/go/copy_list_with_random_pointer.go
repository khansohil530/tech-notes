package main

type RandomListNode struct {
	Val    int
	Next   *RandomListNode
	Random *RandomListNode
}

func copyRandomList(head *RandomListNode) *RandomListNode {
	clone := make(map[*RandomListNode]*RandomListNode)
	curr := head
	for curr != nil {
		clone[curr] = &RandomListNode{Val: curr.Val}
		curr = curr.Next
	}

	curr = head
	cloneCurr, _ := clone[curr]
	for curr != nil {
		cloneCurr.Next, _ = clone[curr.Next]
		cloneCurr.Random, _ = clone[curr.Random]
		curr = curr.Next
		cloneCurr = cloneCurr.Next
	}
	res, _ := clone[head]
	return res
}
