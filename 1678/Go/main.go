package main

import "fmt"

func interpret(command string) string {
	res := ""
	for i := 0; i < len(command); i++ {
		if command[i] == 'G' {
			res += "G"
			continue
		}
		if i < len(command)-1 {
			if command[i] == '(' && command[i+1] == ')' {
				res += "o"
			}
			if command[i] == '(' && command[i+1] == 'a' {
				res += "al"
			}
		}
	}
	return res
}

func main() {
	command := "G()(al)"
	fmt.Println(interpret(command))
	command = "G()()()()(al)"
	fmt.Println(interpret(command))
	command = "(al)G(al)()()G"
	fmt.Println(interpret(command))
}
