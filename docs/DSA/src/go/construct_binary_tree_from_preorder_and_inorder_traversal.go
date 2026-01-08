package main

func buildTree(preorder []int, inorder []int) *TreeNode {
	var helper func(left, right int) *TreeNode
	var preIdx int
	inorderMap := make(map[int]int)
	for idx, v := range inorder {
		inorderMap[v] = idx
	}
	helper = func(left, right int) *TreeNode {
		if left > right {
			return nil
		}

		val := preorder[preIdx]
		preIdx++
		mid := inorderMap[val]
		root := TreeNode{Val: val}
		root.Left = helper(left, mid-1)
		root.Right = helper(mid+1, right)
		return &root
	}

	return helper(0, len(inorder)-1)
}
