package main

func isValid(s string) bool {
	var stk []rune
	op := map[rune]rune{'(': ')', '{': '}', '[': ']'}
	for _, ch := range s {
		val, ok := op[ch]
		if ok {
			stk = append(stk, val)
		} else {
			if len(stk) == 0 || stk[len(stk)-1] != ch {
				return false
			}
			stk = stk[:len(stk)-1]
		}
	}
	return len(stk) == 0
}
