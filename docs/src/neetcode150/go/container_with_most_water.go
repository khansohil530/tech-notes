package main

func maxArea(height []int) int {
	l, r := 0, len(height)-1
	var result int
	for l < r {
		curr := (r - l) * min(height[l], height[r])
		result = max(result, curr)
		if height[r] > height[l] {
			l++
		} else {
			r--
		}
	}
	return result
}
