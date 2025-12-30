package main

import "strconv"

func evalRPN(tokens []string) int {
	var stk []int
	op := map[string]struct{}{
		"+": struct{}{},
		"-": struct{}{},
		"*": struct{}{},
		"/": struct{}{},
	}
	for _, token := range tokens {
		if _, ok := op[token]; !ok {
			token, _ := strconv.Atoi(token)
			stk = append(stk, token)
		} else {
			t2, t1 := stk[len(stk)-1], stk[len(stk)-2]
			var eval int
			if token == "+" {
				eval = t1 + t2
			} else if token == "-" {
				eval = t1 - t2
			} else if token == "*" {
				eval = t1 * t2
			} else if token == "/" {
				eval = t1 / t2
			}
			stk = append(stk[:len(stk)-2], eval)
		}
	}
	return stk[len(stk)-1]
}
