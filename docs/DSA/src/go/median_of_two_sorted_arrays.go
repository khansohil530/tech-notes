package main

import "math"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}

	m, n := len(nums1), len(nums2)
	half := (m + n + 1) / 2

	low, high := 0, m

	for low <= high {
		partA := (low + high) / 2
		partB := half - partA

		var leftA, rightA int
		if partA == 0 {
			leftA = math.MinInt64
		} else {
			leftA = nums1[partA-1]
		}
		if partA == m {
			rightA = math.MaxInt64
		} else {
			rightA = nums1[partA]
		}

		var leftB, rightB int
		if partB == 0 {
			leftB = math.MinInt64
		} else {
			leftB = nums2[partB-1]
		}
		if partB == n {
			rightB = math.MaxInt64
		} else {
			rightB = nums2[partB]
		}

		if leftA <= rightB && leftB <= rightA {
			if (m+n)%2 == 0 {
				return (float64(max(leftA, leftB)) + float64(min(rightA, rightB))) / 2
			}
			return float64(max(leftA, leftB))
		} else if leftA > rightB {
			high = partA - 1
		} else {
			low = partA + 1
		}
	}
	return 0
}
