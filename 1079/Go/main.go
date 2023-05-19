package main

import (
	"fmt"
	"sort"
)

func numTilePossibilities(tiles string) int {
	n := len(tiles)
	used := make([]bool, n)
	cnt := -1
	letters := []byte(tiles)
	sort.SliceStable(letters, func(i, j int) bool {
		return letters[i] < letters[j]
	})
	var backtracking func()
	backtracking = func() {
		cnt++
		for i := 0; i < n; i++ {
			if i > 0 && letters[i] == letters[i-1] && !used[i-1] {
				continue
			}
			if !used[i] {
				used[i] = true
				backtracking()
				used[i] = false
			}
		}
	}
	backtracking()
	return cnt
}
func main() {
	fmt.Println(numTilePossibilities("AAB"))
	fmt.Println(numTilePossibilities("AAABBC"))
	fmt.Println(numTilePossibilities("CDC"))
}
