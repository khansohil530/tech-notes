package main

func dailyTemperatures(temperatures []int) []int {
	result := make([]int, len(temperatures))
	var stk []int
	var l int
	for r, temp := range temperatures {
		for len(stk) != 0 && temperatures[stk[len(stk)-1]] < temp {
			l = stk[len(stk)-1]
			stk = stk[:len(stk)-1]
			result[l] = r - l
		}
		stk = append(stk, r)
	}
	return result
}
