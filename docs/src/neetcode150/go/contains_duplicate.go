package main

func containsDuplicate(nums []int) bool {
	set := make(map[int]interface{})
	for _, num := range nums {
		_, ok := set[num]
		if ok {
			return true
		}
		set[num] = nil
	}
	return false
}
