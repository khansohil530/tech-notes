package main

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	freq := make(map[rune]int)
	for _, ch := range s {
		freq[ch]++
	}

	for _, ch := range t {
		val, ok := freq[ch]
		if !ok || val == 0 {
			return false
		}
		freq[ch]--
	}
	return true
}
