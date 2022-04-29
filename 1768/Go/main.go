package main

import "fmt"

func mergeAlternately(word1 string, word2 string) string {
	res := ""
	m := len(word1)
	n := len(word2)
	if m >= n {
		for i := 0; i < n; i++ {
			res += string(word1[i]) + string(word2[i])
		}
		res += word1[n:]
	} else {
		for i := 0; i < m; i++ {
			res += string(word1[i]) + string(word2[i])
		}
		res += word2[m:]
	}
	return res
}

func main() {
	s1 := "abc"
	s2 := "pqr"
	fmt.Println(mergeAlternately(s1, s2))
	s1 = "ab"
	s2 = "pqrs"
	fmt.Println(mergeAlternately(s1, s2))
}
