package main

import "fmt"

func hammingWeight(num uint32) int {
	s := fmt.Sprintf("%b", num)
	cnt := 0
	for _, v := range s {
		if v == '1' {
			cnt++
		}
	}
	return cnt
}

func main() {
	fmt.Println(hammingWeight(11))
}
