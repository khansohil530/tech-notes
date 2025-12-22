package main

func maxSlidingWindow(nums []int, k int) []int {
	var deque []int
	var result []int
	l := 0
	for r, num := range nums {
		for len(deque) > 0 && nums[deque[len(deque)-1]] < num {
			deque = deque[:len(deque)-1]
		}
		deque = append(deque, r)
		for len(deque) > 0 && deque[0] < l {
			deque = deque[1:]
		}
		if r-l+1 == k {
			result = append(result, nums[deque[0]])
			l++
		}
	}
	return result
}
