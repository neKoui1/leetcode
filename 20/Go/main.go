package main

import "fmt"

func isValid(s string) bool {
	n := len(s)
	myStack := make([]byte, 0)
	pair := make(map[byte]byte)
	pair = map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}
	for i := 0; i < n; i++ {
		if s[i] == ')' || s[i] == ']' || s[i] == '}' {
			if len(myStack) == 0 {
				return false
			}
			if pair[s[i]] != myStack[len(myStack)-1] {
				return false
			}
			myStack = myStack[:len(myStack)-1]
		} else {
			myStack = append(myStack, s[i])
		}
	}

	return len(myStack) == 0
}

func main() {
	fmt.Println(isValid("()"))
	fmt.Println(isValid("()[]{}"))
	fmt.Println(isValid("(]"))
}
