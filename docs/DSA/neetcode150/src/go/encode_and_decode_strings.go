package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"
)

type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	var tokens []string
	for _, str := range strs {
		tokens = append(tokens, fmt.Sprintf("%d#%s", len(str), str))
	}
	return strings.Join(tokens, "")
}

func (s *Solution) Decode(str string) []string {
	var strs []string
	var start, curr int
	for curr < utf8.RuneCountInString(str) {
		for str[curr] != '#' {
			curr++
		}
		size, _ := strconv.Atoi(str[start:curr])
		strs = append(strs, str[curr+1:curr+1+size])
		start = curr + 1 + size
		curr = start
	}
	return strs
}
