package main

import (
	"fmt"
	"sort"
	"strings"
	"unicode"
)

func reorderLogFiles(logs []string) []string {
	sort.SliceStable(logs, func(i, j int) bool {
		m := logs[i]
		n := logs[j]
		s1 := strings.SplitN(m, " ", 2)[1]
		s2 := strings.SplitN(n, " ", 2)[1]
		dig1 := unicode.IsDigit(rune(s1[0]))
		dig2 := unicode.IsDigit(rune(s2[0]))

		if dig1 && dig2 {
			return false
		}
		if !dig1 && !dig2 {
			return s1 < s2 || s1 == s2 && m < n
		}
		return !dig1
	})

	return logs
}
func main() {
	logs := []string{"dig1 8 1 5 1", "let1 art can", "dig2 3 6", "let2 own kit dig", "let3 art zero"}
	fmt.Println(reorderLogFiles(logs))
}
