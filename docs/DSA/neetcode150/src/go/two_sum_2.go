package main

func twoSum2(numbers []int, target int) []int {
	l, r := 0, len(numbers)-1
	for l < r {
		curr := numbers[l] + numbers[r]
		if curr > target {
			r--
		} else if curr < target {
			l++
		} else {
			return []int{l + 1, r + 1}
		}
	}
	return nil
}
