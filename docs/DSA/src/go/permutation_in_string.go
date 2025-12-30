package main

func checkInclusion(s1 string, s2 string) bool {
	counts := make(map[rune]int)
	for _, ch := range s1 {
		counts[ch]++
	}
	var l int
	for r, ch := range s2 {
		counts[ch]--
		for counts[ch] < 0 {
			counts[rune(s2[l])]++
			l++
		}
		if r-l+1 == len(s1) {
			return true
		}
	}
	return false
}
