package main

func productExceptSelf(nums []int) []int {
	result := make([]int, len(nums))
	var pre, post int = 1, 1
	for idx := range nums {
		result[idx] = pre
		pre *= nums[idx]
	}

	for idx := len(nums) - 1; idx >= 0; idx-- {
		result[idx] *= post
		post *= nums[idx]
	}
	return result
}
