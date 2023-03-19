package main

import "fmt"

func vowelStrings(words []string, left int, right int) int {
	res := 0
	for i := left; i <= right; i++ {
		if words[i][0] == 'a' || words[i][0] == 'u' || words[i][0] == 'e' || words[i][0] == 'o' || words[i][0] == 'i' {
			if words[i][len(words[i])-1] == 'a' || words[i][len(words[i])-1] == 'i' || words[i][len(words[i])-1] == 'u' || words[i][len(words[i])-1] == 'e' || words[i][len(words[i])-1] == 'o' {
				res++
			}
		}
	}
	return res
}
func main() {
	fmt.Println(vowelStrings([]string{"are", "amy", "u"}, 0, 2))
	fmt.Println(vowelStrings([]string{"hey", "aeo", "mu", "ooo", "artro"}, 1, 4))
}
