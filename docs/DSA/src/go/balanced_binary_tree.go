package main

import "math"

func isBalanced(root *TreeNode) bool {
	var helper func(node *TreeNode) (bool, int)
	helper = func(node *TreeNode) (bool, int) {
		if node == nil {
			return true, 0
		}

		lb, lh := helper(node.Left)
		rb, rh := helper(node.Right)
		b := math.Abs(float64(lh)-float64(rh)) <= 1 && lb && rb
		h := 1 + max(lh, rh)
		return b, h
	}

	balanced, _ := helper(root)
	return balanced
}
