package main

func calcTime(piles []int, eat int) int {
	var t int
	for _, pile := range piles {
		t += (pile + eat - 1) / eat
	}
	return t
}

func minEatingSpeed(piles []int, h int) int {
	var left, right int
	var sum int
	for _, pile := range piles {
		sum += pile
		right = max(right, pile)
	}
	left = (sum + h - 1) / h

	var mid int
	for left < right {
		mid = left + (right-left)/2
		if calcTime(piles, mid) <= h {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}
