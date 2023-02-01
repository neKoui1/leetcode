package main

import "fmt"

func halvesAreAlike(s string) bool {
	half, cnt := len(s)/2, 0
	for i, v := range s {
		switch v {
		case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
			if i < half {
				cnt++
			} else {
				cnt--
			}
		}
	}
	return cnt == 0
}
func main() {
	fmt.Println(halvesAreAlike("book"))
	fmt.Println(halvesAreAlike("textbook"))
}
