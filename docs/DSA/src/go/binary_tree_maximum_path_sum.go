package main

import "math"

func maxPathSum(root *TreeNode) int {
	var helper func(node *TreeNode) int
	res := math.MinInt
	helper = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		leftMax := max(helper(node.Left), 0)
		rightMax := max(helper(node.Right), 0)
		res = max(res, node.Val+leftMax+rightMax)
		return node.Val + max(leftMax, rightMax)
	}
	helper(root)
	return res
}
