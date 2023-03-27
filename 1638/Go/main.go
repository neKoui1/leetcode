package main

import "fmt"

func countSubstrings(s, t string) (ans int) {
	m, n := len(s), len(t)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			diff := 0
			for k := 0; i+k < m && j+k < n; k++ {
				if s[i+k] != t[j+k] {
					diff++
				}
				if diff > 1 {
					break
				} else if diff == 1 {
					ans++
				}
			}
		}
	}
	return ans
}
func main() {
	fmt.Println(countSubstrings("aba", "baba"))
	fmt.Println(countSubstrings("ab", "bb"))
}
