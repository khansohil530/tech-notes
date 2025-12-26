package main

import "unicode"

func isPalindrome(s string) bool {
	chars := []rune(s)
	l, r := 0, len(s)-1
	for l < r {
		for l < r && !unicode.IsDigit(chars[l]) && !unicode.IsLetter(chars[l]) {
			l++
		}
		for l < r && !unicode.IsDigit(chars[r]) && !unicode.IsLetter(chars[r]) {
			r--
		}
		if unicode.ToLower(chars[l]) != unicode.ToLower(chars[r]) {
			return false
		}
		l++
		r--
	}
	return true
}
