package main

import "fmt"

func reverseLeftWords(s string, n int) string {
	x := []byte(s)
	return string(append(x[n:], x[:n]...))
}
func main() {
	fmt.Println(reverseLeftWords("abcdefg", 2))
	fmt.Println(reverseLeftWords("lrloseumgh", 6))
}
