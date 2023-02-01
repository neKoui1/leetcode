package main

import "fmt"

func findRepeatedDnaSequences(s string) []string {
	res := make([]string, 0)
	temp := make(map[string]int)
	for i := 0; i <= len(s)-10; i++ {
		x := s[i : i+10]
		temp[x]++
		if temp[x] == 2 {
			res = append(res, x)
		}
	}
	return res
}
func main() {
	fmt.Println(findRepeatedDnaSequences("AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT"))
	fmt.Println(findRepeatedDnaSequences("AAAAAAAAAAAAA"))
	fmt.Println(findRepeatedDnaSequences("AAAAAAAAAA"))
	fmt.Println(findRepeatedDnaSequences("AAAAAAAA"))
	fmt.Println(findRepeatedDnaSequences("AAAAAAAAAAA"))
}
