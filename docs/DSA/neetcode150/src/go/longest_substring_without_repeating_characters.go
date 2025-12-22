package main

func lengthOfLongestSubstring(s string) int {
	idxMap := make(map[rune]int)
	var result, l int
	for r, ch := range s {
		if val, ok := idxMap[ch]; ok && val >= l {
			l = val + 1
		}
		idxMap[ch] = r
		result = max(result, r-l+1)
	}
	return result
}
