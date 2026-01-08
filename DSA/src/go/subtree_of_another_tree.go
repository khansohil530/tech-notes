package main

import "strconv"

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	var serialize func(node *TreeNode) string
	serialize = func(node *TreeNode) string {
		if node == nil {
			return "#$"
		}
		return "#" + strconv.Itoa(node.Val) + serialize(node.Left) + serialize(node.Right)
	}

	var z_func func(s string) []int
	z_func = func(s string) []int {
		n := len(s)
		z := make([]int, n)
		l, r := 0, 0
		for i := 1; i < n; i++ {
			if i > r {
				l, r = i, i
				for r < n && s[r] == s[r-l] {
					r++
				}
				z[i] = r - l
				r--
			} else {
				k := i - l
				if z[k] < r-i+1 {
					z[i] = z[k]
				} else {
					l = i
					for r < n && s[r] == s[r-l] {
						r++
					}
					z[i] = r - l
					r--
				}
			}
		}
		return z
	}

	substr, str := serialize(subRoot), serialize(root)

	z := z_func(substr + "|" + str)
	k := len(substr)
	for i := 0; i < len(str); i++ {
		if z[i+k+1] == k {
			return true
		}
	}
	return false
}
