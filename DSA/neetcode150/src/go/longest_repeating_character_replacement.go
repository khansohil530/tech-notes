package main

func characterReplacement(s string, k int) int {
	var maxf, l int
	counts := make(map[rune]int)
	for r, ch := range s {
		counts[ch]++
		maxf = max(maxf, counts[ch])
		if r-l+1-maxf > k { // no more subs
			counts[rune(s[l])]--
			l++
		}
	}
	return len(s) - l
}
