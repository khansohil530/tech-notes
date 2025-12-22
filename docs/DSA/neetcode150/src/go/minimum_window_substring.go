package main

func minWindow(s string, t string) string {
	needs := make(map[rune]int)
	for _, ch := range t {
		needs[ch]++
	}
	missing := len(t)
	var l, R, L int
	for r, ch := range s {
		if val, _ := needs[ch]; val > 0 {
			missing--
		}
		needs[ch]--
		if missing == 0 {
			for l <= r && needs[rune(s[l])] < 0 {
				needs[rune(s[l])]++
				l++
			}
			if (R == 0) || (r-l < R-L) {
				R, L = r+1, l
			}
			needs[rune(s[l])]++
			missing++
			l++
		}
	}
	return s[L:R]
}
