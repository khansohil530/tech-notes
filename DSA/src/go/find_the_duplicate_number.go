package main

func findDuplicate(nums []int) int {
	var slow, fast int
	for fast < len(nums) {
		slow = nums[slow]
		fast = nums[nums[fast]]
		if slow == fast {
			break
		}
	}

	first, second := 0, slow
	for first != second {
		first = nums[first]
		second = nums[second]
	}
	return first
}
