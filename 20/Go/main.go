package main

import "fmt"

func isValid(s string) bool {
	state := 1
	stk := make([]byte, 1)
	for i := 0; i < len(s) && state == 1; {
		switch s[i] {
		case '(':
			stk = append(stk, s[i])
			i++
		case ')':
			if stk != nil && stk[len(stk)-1] == '(' {
				stk = stk[:len(stk)-1]
				i++
			} else {
				state = 0
			}
		}
	}
	if stk == nil && state == 1 {
		return true
	}
	return false
}

func main() {
	fmt.Println(isValid("()"))
	stk := [...]byte{'('}
	fmt.Println(stk)
	fmt.Println(stk[:len(stk)-1])
}
