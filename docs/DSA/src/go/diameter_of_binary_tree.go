package main

func diameterOfBinaryTree(root *TreeNode) int {
	var _helper func(node *TreeNode) (int, int)

	_helper = func(node *TreeNode) (int, int) {
		if node == nil {
			return 0, 0
		}

		leftHeight, leftDiameter := _helper(node.Left)
		rightHeight, rightDiameter := _helper(node.Right)
		height := 1 + max(leftHeight, rightHeight)
		diameter := max(leftDiameter, rightDiameter, leftHeight+rightHeight)
		return height, diameter
	}

	_, diameter := _helper(root)
	return diameter
}
