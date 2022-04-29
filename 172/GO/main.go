package main

import "fmt"

func trailingZeroes(n int) int {
	cnt := 0
	for n/5 != 0 {
		cnt += n / 5
		n /= 5
	}
	return cnt
}

func main() {
	fmt.Println(trailingZeroes(5))
}
