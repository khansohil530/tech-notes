package main

func getIntArray(size int, def int) []int {
	arr := make([]int, size)
	for i := range arr {
		arr[i] = def
	}
	return arr
}

func largestRectangleArea(heights []int) int {
	n := len(heights)
	left := getIntArray(n, -1)
	var stk []int
	for i := 0; i < n; i++ {
		for len(stk) != 0 && heights[i] <= heights[stk[len(stk)-1]] {
			stk = stk[:len(stk)-1]
		}
		if len(stk) != 0 {
			left[i] = stk[len(stk)-1]
		}
		stk = append(stk, i)
	}

	right := getIntArray(n, n)
	stk = []int{}
	for i := n - 1; i >= 0; i-- {
		for len(stk) != 0 && heights[i] <= heights[stk[len(stk)-1]] {
			stk = stk[:len(stk)-1]
		}
		if len(stk) != 0 {
			right[i] = stk[len(stk)-1]
		}
		stk = append(stk, i)
	}

	var result, area int
	for i := range heights {
		area = heights[i] * (right[i] - left[i] - 1)
		result = max(result, area)
	}
	return result
}
