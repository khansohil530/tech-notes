package main

func searchMatrix(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	l, r := 0, m*n
	var mid int
	for l < r {
		mid = l + (r-l)/2
		if matrix[mid/n][mid%n] >= target {
			r = mid
		} else {
			l = mid + 1
		}
	}
	if l < m*n && matrix[l/n][l%n] == target {
		return true
	}
	return false
}
