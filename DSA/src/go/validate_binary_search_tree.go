package main

import "math"

func isValidBST(root *TreeNode) bool {
	var helper func(node *TreeNode, upperLimit, lowerLimit int) bool
	helper = func(node *TreeNode, upperLimit, lowerLimit int) bool {
		if node == nil {
			return true
		}
		return lowerLimit <= node.Val && node.Val <= upperLimit && helper(node.Left, node.Val-1, lowerLimit) &&
			helper(node.Right, upperLimit, node.Val+1)
	}
	return helper(root, math.MaxInt, math.MinInt)
}
