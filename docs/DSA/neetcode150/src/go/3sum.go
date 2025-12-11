package main

import "sort"

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	var result [][]int
	for i := range nums {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		j, k := i+1, len(nums)-1
		for j < k {
			currSum := nums[i] + nums[j] + nums[k]
			if currSum > 0 {
				k--
			} else if currSum < 0 {
				j++
			} else {
				result = append(result, []int{nums[i], nums[j], nums[k]})
				j++
				for j < k && nums[j] == nums[j-1] {
					j++
				}
			}
		}
	}
	return result
}
