package main

import "fmt"

func countAsterisks(s string) int {
	cnt, flag := 0, false
	for _, v := range s {
		if v == '|' {
			flag = !flag
		}
		if !flag {
			if v == '*' {
				cnt++
			}
		}
	}
	return cnt
}
func main() {
	fmt.Println(countAsterisks("l|*e*et|c**o|*de|"))
	fmt.Println(countAsterisks("iamprogrammer"))
	fmt.Println(countAsterisks("yo|uar|e**|b|e***au|tifu|l"))
}
