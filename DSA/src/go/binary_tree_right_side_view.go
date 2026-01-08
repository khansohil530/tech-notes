package main

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	queue := make(chan *TreeNode, 100)
	queue <- root
	var side []int
	var size, i int
	var node *TreeNode
	for len(queue) > 0 {
		size = len(queue)
		for i = 0; i < size; i++ {
			node = <-queue
			if i == size-1 {
				side = append(side, node.Val)
			}
			if node.Left != nil {
				queue <- node.Left
			}
			if node.Right != nil {
				queue <- node.Right
			}
		}
	}
	return side
}
