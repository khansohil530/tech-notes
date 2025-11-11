package main

func twoSum(nums []int, target int) []int {
	idxMap := make(map[int]int)
	for idx, num := range nums {
		if val, ok := idxMap[num]; ok {
			return []int{val, idx}
		}
		idxMap[target-num] = idx
	}
	return []int{}
}
