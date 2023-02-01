package main

import (
	"fmt"
	"math"
)

func minAddToMakeValid(s string) int {
	cnt1, cnt2 := 0, 0
	//cnt1记录(
	//cnt2记录)
	for _, v := range s {
		if cnt1 == 0 && v == ')' {
			cnt2++
			continue
		}
		if v == '(' {
			cnt1++
		} else if v == ')' {
			cnt1--
		}
	}
	return int(math.Abs(float64(cnt1))) + cnt2
}
func main() {
	fmt.Println(minAddToMakeValid("())"))
	fmt.Println(minAddToMakeValid("((("))
}
