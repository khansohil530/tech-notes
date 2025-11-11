package main

func longestConsecutive(nums []int) int {
	set := make(map[int]interface{})
	for _, num := range nums {
		set[num] = nil
	}
	res := 0
	for num, _ := range set {
		if _, ok := set[num-1]; !ok {
			length := 1
			for {
				if _, ok := set[num+length]; ok {
					length++
				} else {
					break
				}
			}
			res = max(res, length)
		}
	}
	return res
}
