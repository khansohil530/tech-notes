package main

func trap(height []int) int {
	l, r := 0, len(height)-1
	maxL, maxR := height[l], height[r]
	var result int
	for l <= r {
		if maxL < maxR {
			maxL = max(maxL, height[l])
			result += maxL - height[l]
			l++
		} else {
			maxR = max(maxR, height[r])
			result += maxR - height[r]
			r--
		}
	}
	return result
}
