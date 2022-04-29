package main

import "fmt"

func uniqueMorseRepresentations(words []string) int {
	morse := []string{
		".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..",
		".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.",
		"...", "-", "..-", "...-", ".--", "-..-", "-.--", "--..",
	}
	line := ""
	res := make([]string, 0)
	for _, v := range words {
		for _, val := range v {
			line += morse[val-97]
		}
		res = append(res, line)
		line = ""
	}
	hash := make(map[string]int)
	for _, v := range res {
		hash[v]++
	}
	return len(hash)
}

func main() {
	words := []string{"gin", "zen", "gig", "msg"}
	fmt.Println(uniqueMorseRepresentations(words))
	words = []string{"a"}
	fmt.Println(uniqueMorseRepresentations(words))
}
