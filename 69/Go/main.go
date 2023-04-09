package main

import "fmt"

func mySqrt(x int) int {
	l, r := 0, x
	for l <= r {
		middle := (l + r) >> 1
		if middle*middle > x {
			r = middle - 1
		} else {
			l = middle + 1
		}
	}
	return l - 1
}
func main() {
	fmt.Println(mySqrt(4))
	fmt.Println(mySqrt(8))
	fmt.Println(mySqrt(1))
	fmt.Println(mySqrt(9))
}
