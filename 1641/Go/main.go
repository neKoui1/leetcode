package main

import "fmt"

func countVowelStrings(n int) int {
	return (n + 4) * (n + 3) * (n + 2) * (n + 1) / 24
}
func main() {
	fmt.Println(countVowelStrings(1))
	fmt.Println(countVowelStrings(2))
	fmt.Println(countVowelStrings(33))
}
