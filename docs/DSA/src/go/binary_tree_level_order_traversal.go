package main

func levelOrder(root *TreeNode) [][]int {
	queue := make(chan *TreeNode, 2000)
	queue <- root
	var levels [][]int
	var level []int
	var node *TreeNode
	var size int
	for len(queue) > 0 {
		level = make([]int, 0)
		size = len(queue)
		for i := 0; i < size; i++ {
			node = <-queue
			if node == nil {
				continue
			}
			level = append(level, node.Val)
			queue <- node.Left
			queue <- node.Right
		}
		if len(level) > 0 {
			levels = append(levels, level)
		}
	}
	return levels
}
