package main

func mergeKLists(lists []*ListNode) *ListNode {
	var l1, l2 *ListNode
	for len(lists) > 1 {
		temp := make([]*ListNode, 0)
		for i := 0; i < len(lists); i += 2 {
			l1 = lists[i]
			l2 = nil
			if i+1 < len(lists) {
				l2 = lists[i+1]
			}
			temp = append(temp, mergeTwoLists(l1, l2))
		}
		lists = temp
	}
	return lists[0]
}
