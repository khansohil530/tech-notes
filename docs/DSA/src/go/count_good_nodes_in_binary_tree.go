package main

func goodNodes(root *TreeNode) int {
	var helper func(node, lastGoodNode *TreeNode) int
	helper = func(node, lastGoodNode *TreeNode) int {
		if node == nil {
			return 0
		}

		tmp := 0
		if node.Val >= lastGoodNode.Val {
			tmp = 1
			lastGoodNode = node
		}
		return tmp + helper(node.Left, lastGoodNode) + helper(node.Right, lastGoodNode)
	}

	return helper(root, root)
}
