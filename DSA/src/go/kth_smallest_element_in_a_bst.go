package main

func kthSmallest(root *TreeNode, k int) int {
	var res int
	var helper func(node *TreeNode)
	helper = func(node *TreeNode) {
		if node == nil {
			return
		}
		helper(node.Left)
		k--
		if k == 0 {
			res = node.Val
			return
		}
		helper(node.Right)
	}
	helper(root)
	return res
}
