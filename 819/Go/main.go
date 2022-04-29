package main

import (
	"fmt"
	"strings"
)

func mostCommonWord(paragraph string, banned []string) string {
	paragraph = strings.ToLower(paragraph)
	i := 0
	ss := make([]string, 0)
	for j, v := range paragraph {
		if v == ' ' {
			ss = append(ss, paragraph[i:j])
			i = j + 1
		} else if v == ',' {
			ss = append(ss, paragraph[i:j])
			i = j + 1
		} else if v == '.' {
			ss = append(ss, paragraph[i:j])
		}
	}
	hash := make(map[string]int)
	max := 0
	for _, v := range ss {
		if v == " " {
			continue
		}
		hash[v]++
	}
	for _, v := range banned {
		delete(hash, v)
	}
	for _, v := range hash {
		if v > max {
			max = v
		}
	}
	for i, v := range hash {
		if v == max {
			return i
		}
	}
	return ""
}
func main() {
	fmt.Println(mostCommonWord("Bob hit a ball, the hit BALL flew far after it was hit.", []string{"hit"}))
	fmt.Println(mostCommonWord("Bob", []string{""}))
}
