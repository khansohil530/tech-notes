package main

func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	var mid int
	for left < right {
		mid = left + (right-left)/2
		if nums[mid] <= nums[right] {
			right = mid
		} else {
			left = mid + 1
		}
	}
	pivot := left
	left, right = 0, len(nums)-1
	if nums[pivot] <= target && target <= nums[right] {
		left = pivot
	} else {
		right = pivot - 1
	}

	for left < right {
		mid = left + (right-left)/2
		if target <= nums[mid] {
			right = mid
		} else {
			left = mid + 1
		}
	}
	if nums[left] == target {
		return left
	}
	return -1
}
