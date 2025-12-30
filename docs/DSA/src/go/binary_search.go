package main

func search(nums []int, target int) int {
	left, right := 0, len(nums)
	var mid int
	for left < right {
		mid = left + (right-left)/2
		if nums[mid] >= target {
			right = mid
		} else {
			left = mid + 1
		}
	}
	if left < len(nums) && nums[left] == target {
		return left
	}
	return -1
}
