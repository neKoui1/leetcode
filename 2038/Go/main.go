package main

import "fmt"

func winnerOfGame(colors string) bool {
	cnt := 0
	for i := 0; i < len(colors)-2; i++ {
		if colors[i:i+3] == "AAA" {
			cnt++
		} else if colors[i:i+3] == "BBB" {
			cnt--
		}
	}
	return cnt > 0
}

func main() {
	fmt.Println(winnerOfGame("AAABABB"))
	fmt.Println(winnerOfGame("AA"))
	fmt.Println(winnerOfGame("ABBBBBBBAAA"))
}
