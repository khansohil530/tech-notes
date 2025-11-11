package main

func getBitmap(s string) [26]int {
	var bitmap [26]int
	for _, ch := range s {
		bitmap[ch-'a']++
	}
	return bitmap
}

func groupAnagrams(strs []string) [][]string {
	freq := make(map[[26]int][]string)
	for _, str := range strs {
		bitmap := getBitmap(str)
		freq[bitmap] = append(freq[bitmap], str)
	}

	var groups [][]string
	for _, group := range freq {
		groups = append(groups, group)
	}
	return groups
}
