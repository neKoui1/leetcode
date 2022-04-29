package main

import "fmt"

func rotateString(s string, goal string) bool {
	origin := s
	temp := s[0]
	s = s[1:] + string(temp)
	for origin != s {
		if s == goal {
			return true
		}
		temp = s[0]
		s = s[1:] + string(temp)
	}
	return false
}

func main() {
	fmt.Println(rotateString("abcde", "cdeab"))
	fmt.Println(rotateString("abcde", "abced"))
}
