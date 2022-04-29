package main

import "fmt"

func Min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

func isAlienSorted(words []string, order string) bool {
	hash := make(map[byte]int)
	for i, v := range order {
		hash[byte(v)] = i
	}

	for i := 0; i < len(words)-1; i++ {
		a := words[i]
		b := words[i+1]
		for j := 0; j < Min(len(a), len(b)); j++ {
			if hash[a[j]] < hash[b[j]] {
				break
			} else if hash[a[j]] > hash[b[j]] {
				return false
			}
			if j == Min(len(a), len(b))-1 && (len(a) > len(b)) {
				return false
			}
		}
	}
	return true
}

func main() {
	words := []string{"hello", "leetcode"}
	order := "hlabcdefgijkmnopqrstuvwxyz"
	fmt.Println(isAlienSorted(words, order))
	words = []string{"word", "world", "row"}
	order = "worldabcefghijkmnpqstuvxyz"
	fmt.Println(isAlienSorted(words, order))
	words = []string{"apple", "app"}
	order = "abcdefghijklmnopqrstuvwxyz"
	fmt.Println(isAlienSorted(words, order))
	words = []string{"hello", "hello"}
	order = "abcdefghijklmnopqrstuvwxyz"
	fmt.Println(isAlienSorted(words, order))
}
